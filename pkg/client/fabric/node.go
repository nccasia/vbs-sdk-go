package fabric

import (
	nodereq "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/node"
	noderes "github.com/nccasia/vbs-sdk-go/pkg/core/model/res/fabric/node"
	"github.com/pkg/errors"
)

const (
	GetTransInfo  = "node/transaction"
	GetBlockInfo  = "node/block"
	GetLedgerInfo = "node/ledger"
)

func (c *FabricClient) GetTransInfo(body nodereq.TransReqDataBody) (*noderes.TransactionResData, error) {
	req := &nodereq.TransReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &noderes.TransactionResData{}

	err := c.Call(GetTransInfo, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", GetTransInfo)
	}
	return res, nil
}

func (c *FabricClient) GetBlockInfo(body nodereq.BlockReqDataBody) (*noderes.BlockResData, error) {
	req := &nodereq.BlockReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &noderes.BlockResData{}

	err := c.Call(GetBlockInfo, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", GetBlockInfo)
	}
	return res, nil
}

func (c *FabricClient) GetLedgerInfo() (*noderes.LedgerResData, error) {
	req := &nodereq.LedgerReqData{}
	req.Header = c.GetHeader()

	res := &noderes.LedgerResData{}

	err := c.Call(GetLedgerInfo, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", GetLedgerInfo)
	}

	return res, nil
}
