package fabric

import (
	"fmt"
	"testing"

	req "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/node"
)

func TestFabricClient_GetTransInfo(t *testing.T) {
	fabricClient := getFabricClient(t)

	tx := req.TransReqDataBody{
		TxId: "99c7091db7548fe00f43f7632f4a2ab4b28655b3ec612f5f288d3af57bbaaa6d",
	}

	res, err := fabricClient.GetTransInfo(tx)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClient_GetBlockInfo(t *testing.T) {
	fabricClient := getFabricClient(t)

	tx := req.BlockReqDataBody{
		BlockNumber: 1,
	}

	res, err := fabricClient.GetBlockInfo(tx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
