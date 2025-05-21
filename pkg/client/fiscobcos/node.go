package fiscobcos

import (
	nodereq "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fiscobcos/node"
	noderes "github.com/nccasia/vbs-sdk-go/pkg/core/model/res/fiscobcos/node"
	"github.com/pkg/errors"
)

const (
	GetTransInfo  = "node/transaction"
	GetBlockInfo  = "node/block"
	GetSystemInfo = "node/system"
)

func (c *FiscoBcosClient) GetTransInfo(body nodereq.TransInfoReqDataBody) (*noderes.TransactionInfoResData, error) {
	req := &nodereq.TransInfoReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &noderes.TransactionInfoResData{}

	err := c.Call(GetTransInfo, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", GetTransInfo)
	}
	return res, nil
}

func (c *FiscoBcosClient) GetBlockInfo(body nodereq.BlockReqDataBody) (*noderes.BlockResData, error) {
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

func (c *FiscoBcosClient) GetSystemInfo() (*noderes.SystemResData, error) {
	req := &nodereq.SystemReqData{}
	req.Header = c.GetHeader()

	res := &noderes.SystemResData{}

	err := c.Call(GetSystemInfo, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", GetSystemInfo)
	}

	return res, nil
}
