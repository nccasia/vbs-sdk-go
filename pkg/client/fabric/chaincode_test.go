package fabric

import (
	"fmt"
	"testing"
	"vbs-sdk-go/pkg/core/config"
	"vbs-sdk-go/pkg/core/fabric/model/req/chaincode"
)

func TestFabricClient_QueryChaincode(t *testing.T) {
	fabricClient := getFabricClient(t)

	args := []string{"37"}
	body := chaincode.QueryChaincodeReqBody{
		UserID:        "tutest10",
		ChaincodeName: "asset-transfer-basic",
		FunctionName:  "ReadAsset",
		Args:          args,
	}
	res, err := fabricClient.QueryChaincode(body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestFabricClient_InvokeChaincode(t *testing.T) {
	fabricClient := getFabricClient(t)

	args := []string{"tu005", "green", "31", "aba", "2000"}
	body := chaincode.InvokeChaincodeReqBody{
		UserID:        "tutest10",
		ChaincodeName: "asset-transfer-basic",
		FunctionName:  "CreateAsset",
		Args:          args,
	}

	res, err := fabricClient.InvokeChaincode(body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func getFabricClient(t *testing.T) *FabricClient {
	api := "http://localhost:8889"
	userCode := "UserCode1"
	appCode := "AppCode1"
	privK := "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgLahbu+lO/KAIF+U7\nps0CHMAyoo+PVtt4GF5T1uXa6GugCgYIKoZIzj0DAQehRANCAATVwUfPrSVZR1/1\nGgfq5pOpohRyPq00Itd2sYCkOJ704yPFIpN7bMcmztQtHJVDh2I+CydXeIyYJ0Tp\ntUQBWUtr\n-----END PRIVATE KEY-----"

	mspDir := "C:\\test"

	config, err := config.NewConfig(api, userCode, appCode, privK, mspDir)
	if err != nil {
		t.Fatal(err)
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err)
	}

	return fabricClient
}
