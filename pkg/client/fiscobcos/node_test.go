package fiscobcos

import (
	"fmt"
	"testing"

	"github.com/nccasia/vbs-sdk-go/pkg/core/config"
	req "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fiscobcos/node"
)

func TestFiscoBcosClientGetTransInfo(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)

	body := req.TransInfoReqDataBody{
		TxHash: "0xc61c3801f299d7795fbeaa80812757f5ec130c121fbf0f591989e8f946af7b8f",
	}

	res, err := fiscobcosClient.GetTransInfo(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFiscoBcosClientGetBlockInfoByBlockNumber(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)

	body := req.BlockReqDataBody{
		BlockNumber: 5,
	}

	res, err := fiscobcosClient.GetBlockInfo(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFiscoBcosClientGetSystemInfo(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)
	res, err := fiscobcosClient.GetSystemInfo()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func getFiscoBcosClient(t *testing.T) *FiscoBcosClient {
	api := "http://localhost:8888"
	userCode := "FiscoBcosUserCode1"
	appCode := "FiscoBcosAppCode1"
	privK := "-----BEGIN PRIVATE KEY-----\nMIGNAgEAMBAGByqGSM49AgEGBSuBBAAKBHYwdAIBAQQg+sSm14WYvHmyr2Boh2nG\nIkcRnDZlsYsv5D4szeoDT+GgBwYFK4EEAAqhRANCAASVHE4ihBmSYo9alPymFHCT\n/MAvnbmqEFS2vNl/1p/n+6CDtTcanLvTcj0fYwNQNl6KlFJtPTJlym16iax7f8Ww\n-----END PRIVATE KEY-----\n"
	mspDir := "C:\\test"

	config, err := config.NewConfig(api, userCode, appCode, privK, mspDir)
	if err != nil {
		t.Fatal(err)
	}

	fiscobcosClient, err := NewFiscoBcosClient(config)

	if err != nil {
		t.Fatal(err)
	}

	return fiscobcosClient
}
