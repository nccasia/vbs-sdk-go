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

	args := []string{"tu002", "green", "31", "aba", "2000"}
	body := chaincode.InvokeChaincodeReqBody{
		UserID:        "tutest9",
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
	privK := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg3iKuG7qbCTPZwF+s
4u2ZDgxxLg/0MhVuQ7XwJsYL6TmhRANCAARfx+NsCeJ+i6+vq/H79B5eUXJ5q2SP
dDs6FB28dUwDR+abx0+C8jWZJ5y17eQUSzmzWGFls8tdbbg39dBCdgas
-----END PRIVATE KEY-----
`

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
