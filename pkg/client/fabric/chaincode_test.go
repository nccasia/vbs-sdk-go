package fabric

import (
	"fmt"
	"testing"

	"github.com/nccasia/vbs-sdk-go/pkg/core/config"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/chaincode"
)

const (
	testChaincodeName = "contract-fabric-testing-2app1748855175839756"
)

func TestFabricClientQueryChaincodeGetAll(t *testing.T) {
	fabricClient := getFabricClient(t)

	args := []string{}
	body := chaincode.QueryChaincodeReq{
		UserID:        "tutest004",
		ChaincodeName: testChaincodeName,
		FunctionName:  "GetAllAssets",
		Args:          args,
	}
	res, err := fabricClient.QueryChaincode(body, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestFabricClientQueryChaincodeGetOne(t *testing.T) {
	fabricClient := getFabricClient(t)

	args := []string{"tu00011"}
	body := chaincode.QueryChaincodeReq{
		UserID:        "tutest003",
		ChaincodeName: testChaincodeName,
		FunctionName:  "ReadAsset",
		Args:          args,
	}
	res, err := fabricClient.QueryChaincode(body, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func getFabricClient(t *testing.T) *FabricClient {
	api := "http://localhost:8889"
	userCode := "UserCode1"
	appCode := "AppCode1"
	privK := "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgxONXM9QezTD7JvSs\ndfMuV64CD8b0jCa2qpc3qJDGjYagCgYIKoZIzj0DAQehRANCAATNAe5f9X2LLSCt\nFP2AFwzYL6dNRb6rckxSMfVd27mjYrKSPelRY/l5bIKLbAi1iXXcUoJie6mwnLdR\nWMl8wJYf\n-----END PRIVATE KEY-----\n"
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

func TestFabricClientPrepareProposal(t *testing.T) {
	fabricClient := getFabricClient(t)

	args := []string{"tu0009", "green", "31", "aba", "20000"}
	body := chaincode.InvokeChaincodeReqBody{
		UserID:        "tutest004",
		ChaincodeName: testChaincodeName,
		FunctionName:  "CreateAsset",
		Args:          args,
	}

	res, err := fabricClient.PrepareProposal(body, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestFabricClientInvokeChaincode(t *testing.T) {
	fabricClient := getFabricClient(t)

	args := []string{"tu00013", "green", "31", "aba", "20000"}
	body := chaincode.InvokeChaincodeReqBody{
		UserID:        "tutest004",
		ChaincodeName: testChaincodeName,
		FunctionName:  "CreateAsset",
		Args:          args,
	}

	res, err := fabricClient.InvokeChaincode(body, nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
