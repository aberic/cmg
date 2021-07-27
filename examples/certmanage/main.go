/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"time"

	"chainmaker.org/chainmaker/pb-go/common"
	"chainmaker.org/chainmaker/pb-go/syscontract"
	sdk "chainmaker.org/chainmaker/sdk-go"
	"chainmaker.org/chainmaker/sdk-go/examples"
)

const (
	sdkConfigOrg1Admin1Path  = "../sdk_configs/sdk_config_org1_admin1.yml"
	sdkConfigOrg1Client1Path = "../sdk_configs/sdk_config_org1_client1.yml"
	sdkConfigOrg2Admin1Path  = "../sdk_configs/sdk_config_org2_admin1.yml"
	sdkConfigOrg3Admin1Path  = "../sdk_configs/sdk_config_org3_admin1.yml"
	sdkConfigOrg4Admin1Path  = "../sdk_configs/sdk_config_org4_admin1.yml"
)

func main() {
	testCertHash()
	testCertManage()
}

func testCertHash() {
	client, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg1Client1Path)
	if err != nil {
		log.Fatalln(err)
	}
	admin1, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg1Admin1Path)
	if err != nil {
		log.Fatalln(err)
	}
	admin2, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg2Admin1Path)
	if err != nil {
		log.Fatalln(err)
	}
	admin3, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg3Admin1Path)
	if err != nil {
		log.Fatalln(err)
	}
	admin4, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg4Admin1Path)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("====================== 用户证书上链 ======================")
	certHash := testCertAdd(client)
	time.Sleep(3 * time.Second)

	fmt.Println("====================== 用户证书查询 ======================")
	certInfos := testQueryCert(client, []string{certHash})
	if len(certInfos.CertInfos) != 1 {
		log.Fatalln("require len(certInfos.CertInfos) == 1")
	}

	fmt.Println("====================== 用户证书删除 ======================")
	testDeleteCert(admin1, admin2, admin3, admin4, []string{certHash})
	time.Sleep(3 * time.Second)

	fmt.Println("====================== 再次查询用户证书 ======================")
	certInfos = testQueryCert(client, []string{certHash})
	if len(certInfos.CertInfos) != 1 {
		log.Fatalln("require len(certInfos.CertInfos) == 1")
	}
	if certInfos.CertInfos[0].Cert != nil {
		log.Fatalln("require certInfos.CertInfos[0].Cert == nil")
	}
}

func testCertManage() {
	// org2 client证书
	var certs = []string{
		"-----BEGIN CERTIFICATE-----\nMIICiDCCAi6gAwIBAgIDCuSTMAoGCCqGSM49BAMCMIGKMQswCQYDVQQGEwJDTjEQ\nMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzEfMB0GA1UEChMWd3gt\nb3JnMi5jaGFpbm1ha2VyLm9yZzESMBAGA1UECxMJcm9vdC1jZXJ0MSIwIAYDVQQD\nExljYS53eC1vcmcyLmNoYWlubWFrZXIub3JnMB4XDTIwMTExNjA2NDYwNFoXDTI1\nMTExNTA2NDYwNFowgZAxCzAJBgNVBAYTAkNOMRAwDgYDVQQIEwdCZWlqaW5nMRAw\nDgYDVQQHEwdCZWlqaW5nMR8wHQYDVQQKExZ3eC1vcmcyLmNoYWlubWFrZXIub3Jn\nMQ8wDQYDVQQLEwZjbGllbnQxKzApBgNVBAMTImNsaWVudDEudGxzLnd4LW9yZzIu\nY2hhaW5tYWtlci5vcmcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQjsmPDqPjx\nikMpRPkmWH8RFgUXwpzwaoMF9OQY6sAty2U8Q6TPlafMbm/xBls//UPZpi5uhwTv\neunkar0HqfvRo3sweTAOBgNVHQ8BAf8EBAMCAaYwDwYDVR0lBAgwBgYEVR0lADAp\nBgNVHQ4EIgQgjqe9Y2WHp+WC/GfKlvwummg3xvKPi9hbDja0QVFKa/EwKwYDVR0j\nBCQwIoAgmZcrtYWpTzN56LDZdqiHah3fG5w0kPaLoEBtyC8GfaEwCgYIKoZIzj0E\nAwIDSAAwRQIgbz8Du0bvtlWVJfBFzUamyfY2OodQDGBbKnr/eFXNeIECIQDnnJs5\nAX2NCT42Be3et+jhwxshehNsYm3WOOdTq/y+yg==\n-----END CERTIFICATE-----\n",
	}

	// org2 client证书的CRL
	var certCrl = "-----BEGIN CRL-----\nMIIBXTCCAQMCAQEwCgYIKoZIzj0EAwIwgYoxCzAJBgNVBAYTAkNOMRAwDgYDVQQI\nEwdCZWlqaW5nMRAwDgYDVQQHEwdCZWlqaW5nMR8wHQYDVQQKExZ3eC1vcmcyLmNo\nYWlubWFrZXIub3JnMRIwEAYDVQQLEwlyb290LWNlcnQxIjAgBgNVBAMTGWNhLnd4\nLW9yZzIuY2hhaW5tYWtlci5vcmcXDTIxMDEyMTA2NDYwM1oXDTIxMDEyMTEwNDYw\nM1owFjAUAgMK5JMXDTI0MDMyMzE1MDMwNVqgLzAtMCsGA1UdIwQkMCKAIJmXK7WF\nqU8zeeiw2Xaoh2od3xucNJD2i6BAbcgvBn2hMAoGCCqGSM49BAMCA0gAMEUCIEgb\nQsHoMkKAKAurOUUfAJpb++DYyxXS3zhvSWPxIUPWAiEAyLSd4TgB9PbSgHyGzS5D\nU1knUTu/4HKTol6GuzmV0Kg=\n-----END CRL-----"

	admin1, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg1Admin1Path)
	if err != nil {
		log.Fatalln(err)
	}
	admin2, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg2Admin1Path)
	if err != nil {
		log.Fatalln(err)
	}
	admin3, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg3Admin1Path)
	if err != nil {
		log.Fatalln(err)
	}
	admin4, err := examples.CreateChainClientWithSDKConf(sdkConfigOrg4Admin1Path)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("====================== 用户证书冻结 ======================")
	testCertManageFrozen(admin1, admin2, admin3, admin4, certs)

	fmt.Println("====================== 用户证书解冻 ======================")
	testCertManageUnfrozen(admin1, admin2, admin3, admin4, certs)

	fmt.Println("====================== 用户证书吊销 ======================")
	testCertManageRevoke(admin1, admin2, admin3, admin4, certCrl)
}

func testCertManageFrozen(admin1, admin2, admin3, admin4 *sdk.ChainClient, certs []string) {
	payload := admin1.CreateCertManageFrozenPayload(certs)

	endorsementEntry1, err := admin1.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry2, err := admin2.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry3, err := admin3.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry4, err := admin4.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := admin1.SendCertManageRequest(payload, []*common.EndorsementEntry{endorsementEntry1, endorsementEntry2, endorsementEntry3, endorsementEntry4}, -1, true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("frozen resp: %+v\n", resp)
}

func testCertManageUnfrozen(admin1, admin2, admin3, admin4 *sdk.ChainClient, certs []string) {
	payload := admin1.CreateCertManageUnfrozenPayload(certs)

	endorsementEntry1, err := admin1.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry2, err := admin2.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry3, err := admin3.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry4, err := admin4.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := admin1.SendCertManageRequest(payload, []*common.EndorsementEntry{endorsementEntry1, endorsementEntry2, endorsementEntry3, endorsementEntry4}, -1, true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("unfrozen resp: %+v\n", resp)
}

func testCertManageRevoke(admin1, admin2, admin3, admin4 *sdk.ChainClient, certCrl string) {
	payload := admin1.CreateCertManageRevocationPayload(certCrl)

	endorsementEntry1, err := admin1.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry2, err := admin2.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry3, err := admin3.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry4, err := admin4.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := admin1.SendCertManageRequest(payload, []*common.EndorsementEntry{endorsementEntry1, endorsementEntry2, endorsementEntry3, endorsementEntry4}, -1, true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("cert revoke resp: %+v\n", resp)
}

func testCertAdd(client *sdk.ChainClient) string {
	resp, err := client.AddCert()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("add cert resp: %+v\n", resp)
	return hex.EncodeToString(resp.ContractResult.Result)
}

func testQueryCert(client *sdk.ChainClient, certHashes []string) *common.CertInfos {
	certInfos, err := client.QueryCert(certHashes)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("query cert resp: %+v\n", certInfos)
	return certInfos
}

func testDeleteCert(admin1, admin2, admin3, admin4 *sdk.ChainClient, certHashes []string) {
	pairs := []*common.KeyValuePair{
		{
			Key:   "cert_hashes",
			Value: []byte(strings.Join(certHashes, "")),
		},
	}

	payload := admin1.CreateCertManagePayload(syscontract.CertManageFunction_CERTS_DELETE.String(), pairs)

	endorsementEntry1, err := admin1.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry2, err := admin2.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry3, err := admin3.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}
	endorsementEntry4, err := admin4.SignCertManagePayload(payload)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := admin1.SendCertManageRequest(payload, []*common.EndorsementEntry{endorsementEntry1, endorsementEntry2, endorsementEntry3, endorsementEntry4}, -1, true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("delete cert resp: %+v\n", resp)
}
