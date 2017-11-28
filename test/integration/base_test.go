/*
Copyright Uhuchain. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
	"os"
	"testing"

	"github.com/uhuchain/uhuchain-api/ledger"
)

var setup ledger.BaseSetupImpl

var initArgs = [][]byte{[]byte("init"), []byte("a"), []byte("100"), []byte("b"), []byte("200")}

func init() {
	log.SetOutput(os.Stdout)
	log.Println("Initialize uhuchain ledger client ")
	setup = ledger.BaseSetupImpl{
		ConfigFile:      os.Getenv("UHU_CONFIG"),
		ChannelID:       os.Getenv("UHU_CHANNELNAME"),
		OrgID:           os.Getenv("UHU_ORG"),
		OrdererOrgID:    os.Getenv("UHU_ORDERER"),
		ChannelConfig:   os.Getenv("UHU_CHANNELCONFIG"),
		ConnectEventHub: true,
		UserID:          os.Getenv("UHU_USERID"),
		AdminUserID:     os.Getenv("UHU_ADMINID"),
	}
}

func TestClient(t *testing.T) {
	client := ledger.FabricClient{}
	client.Init()
	id := "123"
	result, err := client.QueryLedger("automotive", id)
	if err != nil {
		t.Fatalf("Failed to get %s. %s", id, err)
	}
	t.Logf("Got %s", result)
}
