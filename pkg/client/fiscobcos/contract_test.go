package fiscobcos

import (
	"fmt"
	"testing"

	req "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fiscobcos/contract"
)

func TestFiscoBcosClient_QueryContract(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)

	args := []string{}
	body := req.QueryContractReqBody{
		UserID:          "tutest04",
		ContractAddress: "0xb1F8731f406A9Bd4C9f70be544032Ca58C9B2c46",
		FunctionName:    "get",
		Args:            args,
	}
	res, err := fiscobcosClient.QueryContract(body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	fmt.Printf("Body: %s\n", res.Body.Payload)
}

func TestFiscoBcosClient_InvokeContract(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)

	args := []string{"Hello, world"}
	body := req.InvokeContractReqBody{
		UserID:          "tutest04",
		ContractAddress: "0xb1F8731f406A9Bd4C9f70be544032Ca58C9B2c46",
		FunctionName:    "set",
		Args:            args,
	}

	res, err := fiscobcosClient.InvokeContract(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
