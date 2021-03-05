/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chainmaker_sdk_go

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChainClient_CheckNewBlockChainConfig(t *testing.T) {
	client, err := createClient()
	require.Nil(t, err)
	err = client.CheckNewBlockChainConfig()
	require.Nil(t, err)
	fmt.Println("check new block chain config: ok")
}
