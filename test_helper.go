/*
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chainmaker_sdk_go

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/test/bufconn"

	"chainmaker.org/chainmaker/common/v2/ca"
	apipb "chainmaker.org/chainmaker/pb-go/v2/api"
	cmnpb "chainmaker.org/chainmaker/pb-go/v2/common"
	confpb "chainmaker.org/chainmaker/pb-go/v2/config"
	"chainmaker.org/chainmaker/pb-go/v2/syscontract"
	"chainmaker.org/chainmaker/sdk-go/v2/utils"
)

const (
	sdkConfigPathForUT = "./testdata/sdk_config.yml"

	rpcServerTlsCertFile    = "./testdata/crypto-config/wx-org1.chainmaker.org/node/consensus1/consensus1.tls.crt"
	rpcServerTlsPrivKeyFile = "./testdata/crypto-config/wx-org1.chainmaker.org/node/consensus1/consensus1.tls.key"
)

var _ ConnectionPool = (*mockConnectionPool)(nil)

type mockConnectionPool struct {
	connections                    []*networkClient
	logger                         utils.Logger
	userKeyBytes                   []byte
	userCrtBytes                   []byte
	rpcClientMaxReceiveMessageSize int
}

func newMockChainClient(opts ...ChainClientOption) (*ChainClient, error) {
	conf, err := generateConfig(opts...)
	if err != nil {
		return nil, err
	}

	pool, err := newMockConnPool(conf)
	if err != nil {
		return nil, err
	}

	return &ChainClient{
		pool:            pool,
		logger:          conf.logger,
		chainId:         conf.chainId,
		orgId:           conf.orgId,
		userCrtBytes:    conf.userSignCrtBytes,
		userCrt:         conf.userCrt,
		privateKey:      conf.privateKey,
		archiveConfig:   conf.archiveConfig,
		rpcClientConfig: conf.rpcClientConfig,
	}, nil
}

func newMockConnPool(config *ChainClientConfig) (*mockConnectionPool, error) {
	pool := &mockConnectionPool{
		logger:                         config.logger,
		userKeyBytes:                   config.userKeyBytes,
		userCrtBytes:                   config.userCrtBytes,
		rpcClientMaxReceiveMessageSize: config.rpcClientConfig.rpcClientMaxReceiveMessageSize,
	}

	for idx, node := range config.nodeList {
		for i := 0; i < node.connCnt; i++ {
			cli := &networkClient{
				nodeAddr:    node.addr,
				useTLS:      node.useTLS,
				caPaths:     node.caPaths,
				caCerts:     node.caCerts,
				tlsHostName: node.tlsHostName,
				ID:          fmt.Sprintf("%v-%v-%v", idx, node.addr, node.tlsHostName),
			}
			pool.connections = append(pool.connections, cli)
		}
	}

	// 打散，用作负载均衡
	pool.connections = shuffle(pool.connections)

	return pool, nil
}

func (pool *mockConnectionPool) initGRPCConnect(nodeAddr string, useTLS bool, caPaths, caCerts []string,
	tlsHostName string) (*grpc.ClientConn, error) {
	var tlsClient ca.CAClient
	maxCallRecvMsgSize := pool.rpcClientMaxReceiveMessageSize * 1024 * 1024
	if useTLS {
		if len(caCerts) != 0 {
			tlsClient = ca.CAClient{
				ServerName: tlsHostName,
				CaCerts:    caCerts,
				CertBytes:  pool.userCrtBytes,
				KeyBytes:   pool.userKeyBytes,
				Logger:     pool.logger,
			}
		} else {
			tlsClient = ca.CAClient{
				ServerName: tlsHostName,
				CaPaths:    caPaths,
				CertBytes:  pool.userCrtBytes,
				KeyBytes:   pool.userKeyBytes,
				Logger:     pool.logger,
			}
		}

		c, err := tlsClient.GetCredentialsByCA()
		if err != nil {
			return nil, err
		}
		return grpc.Dial("", grpc.WithTransportCredentials(*c),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxCallRecvMsgSize)),
			grpc.WithContextDialer(dialer(useTLS, caPaths, caCerts)))
	}
	return grpc.Dial("", grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxCallRecvMsgSize)),
		grpc.WithContextDialer(dialer(useTLS, caPaths, caCerts)))
}

// 获取空闲的可用客户端连接对象
func (pool *mockConnectionPool) getClient() (*networkClient, error) {
	return pool.getClientWithIgnoreAddrs(nil)
}

func (pool *mockConnectionPool) getClientWithIgnoreAddrs(ignoreAddrs map[string]struct{}) (*networkClient, error) {
	var nc *networkClient

	err := retry.Retry(func(uint) error {
		for _, cli := range pool.connections {

			if ignoreAddrs != nil {
				if _, ok := ignoreAddrs[cli.ID]; ok {
					continue
				}
			}

			if cli.conn == nil || cli.conn.GetState() == connectivity.Shutdown {

				conn, err := pool.initGRPCConnect(cli.nodeAddr, cli.useTLS, cli.caPaths, cli.caCerts, cli.tlsHostName)
				if err != nil {
					pool.logger.Errorf("init grpc connection [nodeAddr:%s] failed, %s", cli.ID, err.Error())
					continue
				}

				cli.conn = conn
				cli.rpcNode = apipb.NewRpcNodeClient(conn)
				nc = cli
				return nil
			}

			s := cli.conn.GetState()
			if s == connectivity.Idle || s == connectivity.Ready {
				nc = cli
				return nil
			}
		}

		return fmt.Errorf("all client connections are busy")

	}, strategy.Wait(retryInterval*time.Millisecond), strategy.Limit(retryLimit))

	if err != nil {
		return nil, err
	}

	return nc, nil
}

func (pool *mockConnectionPool) getLogger() utils.Logger {
	return pool.logger
}

// Close 关闭连接池
func (pool *mockConnectionPool) Close() error {
	for _, c := range pool.connections {
		if c.conn == nil {
			continue
		}

		if err := c.conn.Close(); err != nil {
			pool.logger.Errorf("stop %s connection failed, %s",
				c.nodeAddr, err.Error())

			continue
		}
	}

	return nil
}

type mockRpcNodeServer struct {
	apipb.UnimplementedRpcNodeServer
}

func (s *mockRpcNodeServer) SendRequest(ctx context.Context, req *cmnpb.TxRequest) (*cmnpb.TxResponse, error) {
	switch req.Payload.TxType {
	case cmnpb.TxType_ARCHIVE:
		switch req.Payload.Method {

		case syscontract.ArchiveFunction_ARCHIVE_BLOCK.String():
			return &cmnpb.TxResponse{Code: cmnpb.TxStatusCode_SUCCESS}, nil
		case syscontract.ArchiveFunction_RESTORE_BLOCK.String():
			return &cmnpb.TxResponse{Code: cmnpb.TxStatusCode_SUCCESS}, nil
		default:
			return &cmnpb.TxResponse{Code: cmnpb.TxStatusCode_CONTRACT_FAIL}, nil
		}
	}
	return &cmnpb.TxResponse{}, nil
}

func (s *mockRpcNodeServer) Subscribe(req *cmnpb.TxRequest, server apipb.RpcNode_SubscribeServer) error {
	//var (
	//	errCode cmnerr.ErrCode
	//	errMsg  string
	//)

	//tx := &cmnpb.Transaction{
	//	Header:           req.Header,
	//	RequestPayload:   req.Payload,
	//	RequestSignature: req.Signature,
	//	Result:           nil,
	//}

	//errCode, errMsg = s.validate(tx)
	//if errCode != cmnerr.ERR_CODE_OK {
	//	return status.Error(codes.Unauthenticated, errMsg)
	//}

	switch req.Payload.TxType {
	case cmnpb.TxType_SUBSCRIBE:
	}

	return nil
}

func (s *mockRpcNodeServer) GetChainMakerVersion(ctx context.Context,
	req *confpb.ChainMakerVersionRequest) (*confpb.ChainMakerVersionResponse, error) {
	return &confpb.ChainMakerVersionResponse{
		Code:    0,
		Message: "OK",
		Version: "1.0.0",
	}, nil
}

func (s *mockRpcNodeServer) CheckNewBlockChainConfig(ctx context.Context,
	req *confpb.CheckNewBlockChainConfigRequest) (*confpb.CheckNewBlockChainConfigResponse, error) {
	return &confpb.CheckNewBlockChainConfigResponse{
		Code: 0,
	}, nil
}

func dialer(useTLS bool, caPaths, caCerts []string) func(context.Context, string) (net.Conn, error) {
	var opts []grpc.ServerOption
	var tlsRPCServer ca.CAServer

	if useTLS {
		if len(caCerts) != 0 {
			tlsRPCServer = ca.CAServer{
				CaCerts:  caCerts,
				CertFile: rpcServerTlsCertFile,
				KeyFile:  rpcServerTlsPrivKeyFile,
			}
		} else {
			tlsRPCServer = ca.CAServer{
				CaPaths:  caPaths,
				CertFile: rpcServerTlsCertFile,
				KeyFile:  rpcServerTlsPrivKeyFile,
			}
		}

		c, err := tlsRPCServer.GetCredentialsByCA(true)
		if err != nil {
			log.Fatalf("new gRPC failed, GetTLSCredentialsByCA err: %v\n", err)
		}

		opts = append(opts, grpc.Creds(*c))
	}

	server := grpc.NewServer(opts...)
	listener := bufconn.Listen(1024 * 1024)

	apipb.RegisterRpcNodeServer(server, &mockRpcNodeServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
