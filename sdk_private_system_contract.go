/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chainmaker_sdk_go

import (
	"chainmaker.org/chainmaker/pb-go/common"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
)

func (cc *ChainClient) SaveDir(orderId, txId string,
	privateDir *common.StrSlice, withSyncResult bool, timeout int64) (*common.TxResponse, error) {

	if txId == "" {
		txId = GetRandTxId()
	}

	cc.logger.Infof("[SDK] begin to save dir , [contract:%s]/[method:%s]/[txId:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_DIR.String(),
		txId,
	)

	// 构造Payload
	priDirBytes, err := privateDir.Marshal()
	if err != nil {
		return nil, fmt.Errorf("serielized private dir failed, %s", err.Error())
	}

	pairs := paramsMap2KVPairs(map[string]string{
		"order_id":    orderId,
		"private_dir": string(priDirBytes),
	})

	payloadBytes, err := constructSystemContractPayload(
		cc.chainId,
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_DIR.String(),
		pairs,
		defaultSequence,
	)
	if err != nil {
		return nil, fmt.Errorf("construct save dir payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequestWithTimeout(common.TxType_INVOKE_SYSTEM_CONTRACT, txId, payloadBytes, timeout)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.ContractResult = &common.ContractResult{
				Code:    common.ContractResultCode_OK,
				Message: common.ContractResultCode_OK.String(),
				Result:  []byte(txId),
			}
		} else {
			contractResult, err := cc.getSyncResult(txId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != common.ContractResultCode_OK {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			}

			resp.ContractResult = contractResult
		}
	}

	if resp.Code != common.TxStatusCode_SUCCESS || resp.Message != "OK" {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp, nil
}

func (cc *ChainClient) GetContract(contractName, codeHash string) (*common.PrivateGetContract, error) {

	cc.logger.Infof("[SDK] begin to get contract , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_CONTRACT.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"contract_name": contractName,
		"code_hash":     codeHash,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_CONTRACT.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get contract payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	contractInfo := &common.PrivateGetContract{}
	if err = proto.Unmarshal(resp.ContractResult.Result, contractInfo); err != nil {
		return nil, fmt.Errorf("GetContract unmarshal contract info payload failed, %s", err.Error())
	}

	return contractInfo, nil
}

func (cc *ChainClient) SaveData(contractName string, contractVersion string, isDeployment bool, codeHash []byte,
	reportHash []byte, result *common.ContractResult, codeHeader []byte, txId string, rwSet *common.TxRWSet,
	sign []byte, events *common.StrSlice, privateReq []byte, withSyncResult bool,
	timeout int64) (*common.TxResponse, error) {
	if txId == "" {
		txId = GetRandTxId()
	}

	cc.logger.Infof("[SDK] begin to save data , [contract:%s]/[method:%s]/[txId:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_DATA.String(),
		txId,
	)

	// 构造Payload
	var rwSetStr string
	if rwSet != nil {
		rwb, err := rwSet.Marshal()
		if err != nil {
			return nil, fmt.Errorf("construct save data payload failed, %s", err.Error())
		}
		rwSetStr = string(rwb)
	}

	var eventsStr string
	if events != nil {
		eb, err := events.Marshal()
		if err != nil {
			return nil, fmt.Errorf("construct save data payload failed, %s", err.Error())
		}
		eventsStr = string(eb)
	}

	var resultStr string
	if result != nil {
		result, err := result.Marshal()
		if err != nil {
			return nil, fmt.Errorf("construct save data payload failed, %s", err.Error())
		}
		resultStr = string(result)
	}

	deployStr := strconv.FormatBool(isDeployment)
	pairsMap := map[string]string {
		"result":        resultStr,
		"code_header":   string(codeHeader),
		"contract_name": contractName,
		"version":       contractVersion,
		"is_deploy":     deployStr,
		"code_hash":     string(codeHash),
		"rw_set":        rwSetStr,
		"events":        eventsStr,
		"report_hash":   string(reportHash),
		"sign":          string(sign),
	}
	if isDeployment {
		pairsMap["deploy_req"] = string(privateReq)
	} else {
		pairsMap["private_req"] = string(privateReq)
	}
	pairs := paramsMap2KVPairs(pairsMap)

	payloadBytes, err := constructSystemContractPayload(
		cc.chainId,
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_DATA.String(),
		pairs,
		defaultSequence,
	)
	if err != nil {
		return nil, fmt.Errorf("construct save data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequestWithTimeout(common.TxType_INVOKE_SYSTEM_CONTRACT, txId, payloadBytes, timeout)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.ContractResult = &common.ContractResult{
				Code:    common.ContractResultCode_OK,
				Message: common.ContractResultCode_OK.String(),
				Result:  []byte(txId),
			}
		} else {
			contractResult, err := cc.getSyncResult(txId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != common.ContractResultCode_OK {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			}

			resp.ContractResult = contractResult
		}
	}

	if resp.Code != common.TxStatusCode_SUCCESS || resp.Message != "OK" {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), resp.Message)
	}

	return resp, nil
}

func (cc *ChainClient) GetData(contractName, key string) ([]byte, error) {
	cc.logger.Infof("[SDK] begin to get data , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_DATA.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"contract_name": contractName,
		"key":           key,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_DATA.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) GetDir(orderId string) ([]byte, error) {
	cc.logger.Infof("[SDK] begin to get data , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_DIR.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"order_id": orderId,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_DIR.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) CheckCallerCertAuth(payload string, orgIds []string, signPairs []*common.SignInfo) (
	*common.TxResponse, error) {
	cc.logger.Infof("[SDK] begin to check caller cert auth  , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_CHECK_CALLER_CERT_AUTH.String(),
	)

	orgIdsJson, err := json.Marshal(orgIds)
	if err != nil {
		return nil, fmt.Errorf("json marshal orgIds failed, err: %v", err)
	}
	signPairsJson, err := json.Marshal(signPairs)
	if err != nil {
		return nil, fmt.Errorf("json marshal signPairs failed, err: %v", err)
	}
	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"payload":   payload,
		"org_ids": 	string(orgIdsJson),
		"sign_pairs": string(signPairsJson),
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_CHECK_CALLER_CERT_AUTH.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp, nil
}

//
//func (cc *ChainClient) CheckCallerCertAuth(privateComputeRequest string) (*common.TxResponse, error) {
//	cc.logger.Infof("[SDK] begin to check caller cert auth  , [contract:%s]/[method:%s]",
//		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
//		common.PrivateComputeContractFunction_CHECK_CALLER_CERT_AUTH.String(),
//	)
//
//	// 构造Payload
//	pairs := paramsMap2KVPairs(map[string]string{
//		"private_req":   privateComputeRequest,
//	})
//
//	payloadBytes, err := constructQueryPayload(
//		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
//		common.PrivateComputeContractFunction_CHECK_CALLER_CERT_AUTH.String(),
//		pairs,
//	)
//	if err != nil {
//		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
//	}
//
//	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
//	if err != nil {
//		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
//	}
//
//	if err = checkProposalRequestResp(resp, true); err != nil {
//		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
//	}
//
//	return resp, nil
//}

func (cc *ChainClient) SaveEnclaveCACert(enclaveCACert, txId string, withSyncResult bool, timeout int64) (*common.TxResponse, error) {
	if txId == "" {
		txId = GetRandTxId()
	}

	cc.logger.Infof("[SDK] begin to save ca cert , [contract:%s]/[method:%s]/[txId:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_CA_CERT.String(),
		txId,
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"ca_cert": enclaveCACert,
	})

	payloadBytes, err := constructSystemContractPayload(
		cc.chainId,
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_CA_CERT.String(),
		pairs,
		defaultSequence,
	)
	if err != nil {
		return nil, fmt.Errorf("construct save cert payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequestWithTimeout(common.TxType_INVOKE_SYSTEM_CONTRACT, txId, payloadBytes, timeout)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.ContractResult = &common.ContractResult{
				Code:    common.ContractResultCode_OK,
				Message: common.ContractResultCode_OK.String(),
				Result:  []byte(txId),
			}
		} else {
			contractResult, err := cc.getSyncResult(txId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != common.ContractResultCode_OK {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			} else {
				resp.ContractResult = contractResult
			}
		}
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp, nil
}

func (cc *ChainClient) GetEnclaveCACert() ([]byte, error) {
	cc.logger.Infof("[SDK] begin to get ca cert , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_CA_CERT.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_CA_CERT.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) SaveEnclaveReport(enclaveId, report, txId string, withSyncResult bool, timeout int64) (*common.TxResponse, error) {
	if txId == "" {
		txId = GetRandTxId()
	}

	cc.logger.Infof("[SDK] begin to save enclave report , [contract:%s]/[method:%s]/[txId:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_ENCLAVE_REPORT.String(),
		txId,
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"enclave_id": enclaveId,
		"report":     report,
	})

	payloadBytes, err := constructSystemContractPayload(
		cc.chainId,
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_ENCLAVE_REPORT.String(),
		pairs,
		defaultSequence,
	)
	if err != nil {
		return nil, fmt.Errorf("construct save cert payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequestWithTimeout(common.TxType_INVOKE_SYSTEM_CONTRACT, txId, payloadBytes, timeout)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.ContractResult = &common.ContractResult{
				Code:    common.ContractResultCode_OK,
				Message: common.ContractResultCode_OK.String(),
				Result:  []byte(txId),
			}
		} else {
			contractResult, err := cc.getSyncResult(txId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != common.ContractResultCode_OK {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			} else {
				resp.ContractResult = contractResult
			}
		}
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp, nil
}

func (cc *ChainClient) SaveRemoteAttestationProof(proof, txId string, withSyncResult bool, timeout int64) (*common.TxResponse, error) {
	if txId == "" {
		txId = GetRandTxId()
	}

	cc.logger.Infof("[SDK] begin to save_remote_attestation_proof , [contract:%s]/[method:%s]/[txId:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_REMOTE_ATTESTATION.String(),
		txId,
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"proof": proof,
	})

	payloadBytes, err := constructSystemContractPayload(
		cc.chainId,
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_REMOTE_ATTESTATION.String(),
		pairs,
		defaultSequence,
	)
	if err != nil {
		return nil, fmt.Errorf("construct save quote  payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequestWithTimeout(common.TxType_INVOKE_SYSTEM_CONTRACT, txId, payloadBytes, timeout)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.ContractResult = &common.ContractResult{
				Code:    common.ContractResultCode_OK,
				Message: common.ContractResultCode_OK.String(),
				Result:  []byte(txId),
			}
		} else {
			contractResult, err := cc.getSyncResult(txId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != common.ContractResultCode_OK {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			} else {
				resp.ContractResult = contractResult
			}
		}
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp, nil
}

func (cc *ChainClient) GetEnclaveEncryptPubKey(enclaveId string) ([]byte, error) {
	cc.logger.Infof("[SDK] begin get_enclave_encrypt_pub_key() , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_ENCRYPT_PUB_KEY.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"enclave_id": enclaveId,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_ENCRYPT_PUB_KEY.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) GetEnclaveVerificationPubKey(enclaveId string) ([]byte, error) {
	cc.logger.Infof("[SDK] begin to get data , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_VERIFICATION_PUB_KEY.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"enclave_id": enclaveId,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_VERIFICATION_PUB_KEY.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) GetEnclaveReport(enclaveId string) ([]byte, error) {
	cc.logger.Infof("[SDK] begin to get data , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_REPORT.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"enclave_id": enclaveId,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_REPORT.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) GetEnclaveChallenge(enclaveId string) ([]byte, error) {
	cc.logger.Infof("[SDK] begin to get data , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_CHALLENGE.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"enclave_id": enclaveId,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_CHALLENGE.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) GetEnclaveSignature(enclaveId string) ([]byte, error) {
	cc.logger.Infof("[SDK] begin to get data , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_SIGNATURE.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"enclave_id": enclaveId,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_SIGNATURE.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) GetEnclaveProof(enclaveId string) ([]byte, error) {
	cc.logger.Infof("[SDK] begin to get data , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_VERIFICATION_PUB_KEY.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"enclave_id": enclaveId,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_ENCLAVE_PROOF.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}