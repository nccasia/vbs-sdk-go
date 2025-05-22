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
		ContractAddress: "0xC8eAd4B26b2c6Ac14c9fD90d9684c9Bc2cC40085",
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

	args := []string{"Hello, World!!!!!!!!"}
	body := req.InvokeContractReqBody{
		UserID:          "tutest04",
		ContractAddress: "0xC8eAd4B26b2c6Ac14c9fD90d9684c9Bc2cC40085",
		FunctionName:    "set",
		Args:            args,
	}

	res, err := fiscobcosClient.InvokeContract(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
