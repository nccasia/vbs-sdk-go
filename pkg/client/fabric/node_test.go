package fabric

import (
	"fmt"
	"testing"

	req "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/node"
)

func TestFabricClientGetTransInfo(t *testing.T) {
	fabricClient := getFabricClient(t)

	tx := req.TransReqDataBody{
		TxId: "93833649b68e9d939f9ead9ae63af4ce853bad78f4302f8f07f50f78e80f3518",
	}

	res, err := fabricClient.GetTransInfo(tx)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClientGetBlockInfoByBlockNumber(t *testing.T) {
	fabricClient := getFabricClient(t)

	tx := req.BlockReqDataBody{
		BlockNumber: 32,
	}

	res, err := fabricClient.GetBlockInfo(tx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClientGetBlockInfoBlockHash(t *testing.T) {
	fabricClient := getFabricClient(t)

	tx := req.BlockReqDataBody{
		BlockHash: "b7fd57b34198a9a3a58617e2cf1da02d9ed9185ca37e77b969998c4ee925d2e3",
	}

	res, err := fabricClient.GetBlockInfo(tx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFabricClientGetLedgerInfo(t *testing.T) {
	fabricClient := getFabricClient(t)
	res, err := fabricClient.GetLedgerInfo()

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
