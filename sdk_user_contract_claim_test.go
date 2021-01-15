/**
 * @Author: jasonruan
 * @Date:   2021-01-15 12:28:45
 **/
package chainmaker_sdk_go

import (
	"chainmaker.org/chainmaker-go/chainmaker-sdk-go/pb"
	"chainmaker.org/chainmaker-go/common/random/uuid"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	claimContractName = "claim001"
	claimVersion = "1.0.0"
	claimByteCodePath = "./testdata/claim-wasm-demo/fact-rust-0.7.2.wasm"
)

func TestUserContractClaim(t *testing.T) {
	client, err := createClientWithConfig()
	require.Nil(t, err)

	admin1, err := createAdmin(orgId1)
	require.Nil(t, err)
	admin2, err := createAdmin(orgId2)
	require.Nil(t, err)
	admin3, err := createAdmin(orgId3)
	require.Nil(t, err)
	admin4, err := createAdmin(orgId4)
	require.Nil(t, err)

	fmt.Println("====================== 创建合约 ======================")
	testUserContractClaimCreate(t, client, admin1, admin2, admin3, admin4, true)

	fmt.Println("====================== 调用合约 ======================")
	fileHash, err := testUserContractClaimInvoke(t, client, "save", true)
	require.Nil(t, err)

	fmt.Println("====================== 执行合约查询接口 ======================")
	params := map[string]string {
		"file_hash": fileHash,
	}
	testUserContractClaimQuery(t, client, "find_by_file_hash", params)
}

func testUserContractClaimCreate(t *testing.T, client *ChainClient,
	admin1, admin2, admin3, admin4 *ChainClient, withSyncResult bool) {


	resp, err := createUserContract(client, admin1, admin2, admin3, admin4,
		claimContractName, claimVersion, claimByteCodePath, pb.RuntimeType_WASMER, []*pb.KeyValuePair{}, withSyncResult)
	require.Nil(t, err)

	fmt.Printf("CREATE claim contract resp: %+v\n", resp)
}

func testUserContractClaimInvoke(t *testing.T, client *ChainClient,
	method string, withSyncResult bool) (string, error) {

	curTime := fmt.Sprintf("%d", CurrentTimeMillisSeconds())
	fileHash := uuid.GetUUID()
	params := map[string]string {
		"time": curTime,
		"file_hash": fileHash,
		"file_name": fmt.Sprintf("file_%s", curTime),
	}

	err := invokeUserContract(client, claimContractName, method, "", params, withSyncResult)
	if err != nil {
		return "", err
	}

	return fileHash, nil
}

func testUserContractClaimQuery(t *testing.T, client *ChainClient,
	method string, params map[string]string) {
	resp, err := client.QueryContract(claimContractName, method, params, -1)
	require.Nil(t, err)
	fmt.Printf("QUERY claim contract resp: %+v\n", resp)
}
