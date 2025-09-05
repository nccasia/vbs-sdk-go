package fabric

import (
	"fmt"
	"testing"

	req "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/event"
)

func TestFabricClientEventRegister(t *testing.T) {
	fabricClient := getFabricClient(t)

	body := req.RegisterEventReqBody{
		ChainCode:  "basic-transfer-contractapp1747897523507630",
		UserID:     "tutest001",
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

func TestFabricClientBlockEventRegister(t *testing.T) {
	fabricClient := getFabricClient(t)

	body := req.RegisterEventReqBody{
		ChainCode:  "basic-transfer-contractapp1747897523507630",
		UserID:     "tutest001",
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

func TestFabricClientEventQuery(t *testing.T) {
	fabricClient := getFabricClient(t)

	body := req.QueryEventReqBody{
		UserID: "tutest001",
	}
	res, err := fabricClient.EventQuery(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClientEventRemove(t *testing.T) {
	fabricClient := getFabricClient(t)
	body := req.RemoveEventReqBody{
		EventId: "9b7e0900-f671-4428-aec5-62889a110aa3",
		UserID:  "tutest001",
	}

	res, err := fabricClient.EventRemove(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
