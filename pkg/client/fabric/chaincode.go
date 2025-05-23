package fabric

import (
	ccreq "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/chaincode"
	ccres "github.com/nccasia/vbs-sdk-go/pkg/core/model/res/fabric/chaincode"

	"github.com/pkg/errors"
)

const (
	QueryChaincode  = "chaincode/query"
	InvokeChaincode = "chaincode/invoke"
)

func (c *FabricClient) QueryChaincode(body ccreq.QueryChaincodeReqBody) (*ccres.QueryChaincodeResData, error) {
	req := &ccreq.QueryChaincodeReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &ccres.QueryChaincodeResData{}

	err := c.Call(QueryChaincode, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", QueryChaincode)
	}

	return res, nil
}

func (c *FabricClient) InvokeChaincode(body ccreq.InvokeChaincodeReqBody) (*ccres.InvokeChaincodeResData, error) {
	req := &ccreq.InvokeChaincodeReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &ccres.InvokeChaincodeResData{}

	err := c.Call(InvokeChaincode, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", InvokeChaincode)
	}

	return res, nil
}
