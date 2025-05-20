package fiscobcos

import (
	ctreq "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fiscobcos/contract"
	ctres "github.com/nccasia/vbs-sdk-go/pkg/core/model/res/fiscobcos/contract"
	"github.com/pkg/errors"
)

const (
	QueryContract  = "contracts/query"
	InvokeContract = "contracts/invoke"
)

func (c *FiscoBcosClient) QueryContract(body ctreq.QueryContractReqBody) (*ctres.QueryContractResData, error) {
	req := &ctreq.QueryContractReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &ctres.QueryContractResData{}

	err := c.Call(QueryContract, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", QueryContract)
	}

	return res, nil
}

func (c *FiscoBcosClient) InvokeContract(body ctreq.InvokeContractReqBody) (*ctres.InvokeContractResData, error) {
	req := &ctreq.InvokeContractReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &ctres.InvokeContractResData{}

	err := c.Call(InvokeContract, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", InvokeContract)
	}

	return res, nil
}
