package fabric

import (
	"fmt"
	"testing"

	req "vbs-sdk-go/pkg/core/fabric/model/req/event"
)

func TestFabricClient_EventRegister(t *testing.T) {
	fabricClient := getFabricClient(t)

	body := req.RegisterEventReqBody{
		ChainCode:  "asset-transfer-basic",
		UserID:     "tutest8",
		EventKey:   "TransferAsset",
		NotifyUrl:  "http://localhost:8888/api/event/chaincode-hook",
		AttachArgs: "",
	}

	res, err := fabricClient.EventRegister(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClient_BlockEventRegister(t *testing.T) {
	fabricClient := getFabricClient(t)

	body := req.RegisterEventReqBody{
		ChainCode:  "asset-transfer-basic",
		UserID:     "tutest9",
		EventKey:   "",
		NotifyUrl:  "http://localhost:8888/api/event/block-hook",
		AttachArgs: "",
	}

	res, err := fabricClient.BlockEventRegister(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClient_EventQuery(t *testing.T) {
	fabricClient := getFabricClient(t)

	body := req.QueryEventReqBody{
		UserID: "tutest9",
	}
	res, err := fabricClient.EventQuery(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClient_EventRemove(t *testing.T) {
	fabricClient := getFabricClient(t)
	body := req.RemoveEventReqBody{
		EventId: "c5ec763c-f852-4c11-821b-8702b03db2fb",
		UserID:  "tutest9",
	}

	res, err := fabricClient.EventRemove(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
