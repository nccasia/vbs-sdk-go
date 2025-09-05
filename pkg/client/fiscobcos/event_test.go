package fiscobcos

import (
	"fmt"
	"testing"

	req "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fiscobcos/event"
)

func TestFiscoBcosClientEventRegister(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)

	body := req.RegisterEventReqBody{
		ContractAddress: "0xC8eAd4B26b2c6Ac14c9fD90d9684c9Bc2cC40085",
		UserID:          "tutest04",
		EventName:       "setValue",
		NotifyUrl:       "http://localhost:8888/api/event/chaincode-hook",
		AttachArgs:      "",
	}

	res, err := fiscobcosClient.EventRegister(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFiscoBcosClientBlockEventRegister(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)

	body := req.RegisterEventReqBody{
		ContractAddress: "",
		UserID:          "tutest04",
		EventName:       "",
		NotifyUrl:       "http://localhost:8888/api/event/block-hook",
		AttachArgs:      "",
	}

	res, err := fiscobcosClient.BlockEventRegister(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFiscoBcosClientEventQuery(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)

	body := req.QueryEventReqBody{
		UserID: "tutest04",
	}
	res, err := fiscobcosClient.EventQuery(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFiscoBcosClientEventRemove(t *testing.T) {
	fiscobcosClient := getFiscoBcosClient(t)
	body := req.RemoveEventReqBody{
		EventId: "1249c5c1-0e62-485e-8116-d9e81694794c",
		UserID:  "tutest04",
	}

	res, err := fiscobcosClient.EventRemove(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
