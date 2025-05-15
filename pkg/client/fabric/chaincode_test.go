package fabric

import (
	"fmt"
	"testing"

	"github.com/nccasia/vbs-sdk-go/pkg/core/config"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/chaincode"
)

func TestFabricClient_QueryChaincode_GetAll(t *testing.T) {
	fabricClient := getFabricClient(t)

	args := []string{}
	body := chaincode.QueryChaincodeReqBody{
		UserID:        "tutest01",
		ChaincodeName: "contract-testingapp1747220033738860",
		FunctionName:  "GetAllAssets",
		Args:          args,
	}
	res, err := fabricClient.QueryChaincode(body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestFabricClient_QueryChaincode_GetOne(t *testing.T) {
	fabricClient := getFabricClient(t)

	args := []string{"tu001"}
	body := chaincode.QueryChaincodeReqBody{
		UserID:        "tutest01",
		ChaincodeName: "contract-testingapp1747220033738860",
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

	args := []string{"tu001", "green", "31", "aba", "2000"}
	body := chaincode.InvokeChaincodeReqBody{
		UserID:        "tutest01",
		ChaincodeName: "contract-testingapp1747220033738860",
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
