package fabric

import (
	"fmt"

	ccreq "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/chaincode"
	ccres "github.com/nccasia/vbs-sdk-go/pkg/core/model/res/fabric/chaincode"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/trans/fabric/proposal"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/userdata"

	"github.com/pkg/errors"
)

const (
	QueryChaincode  = "chaincode/query"
	InvokeChaincode = "chaincode/invoke"
	PrepareProposal = "chaincode/proposal/prepare"
	SubmitProposal  = "chaincode/proposal/submit"

	// Error message constants
	ErrUserLoadFailed               = "user [%s] load failed"
	ErrFailedToCreateSignedProposal = "failed to create signed proposal"
	ErrCallHasError                 = "call %s has error"
)

func (c *FabricClient) QueryChaincode(body ccreq.QueryChaincodeReq, user *userdata.UserData) (*ccres.QueryChaincodeResData, error) {
	var err error
	if user == nil {
		user, err = c.LoadUser(body.UserID)
		if err != nil {
			return nil, errors.WithMessagef(err, ErrUserLoadFailed, body.UserID)
		}
	}

	channelId := c.GetAppInfo().GetChannelId()
	fmt.Printf("MspId: %s\n", user.MspId)
	fmt.Printf("ChannelId: %s\n", channelId)

	signedProposal, err := proposal.CreateSignedProposal(channelId, body.ChaincodeName, body.FunctionName, body.Args, user)
	if err != nil {
		return nil, errors.WithMessagef(err, ErrFailedToCreateSignedProposal)
	}
	fmt.Printf("TxId: %s\n", signedProposal.TxId)

	req := &ccreq.QueryChaincodeReqData{}
	req.Header = c.GetHeader()
	resBody := ccreq.SignedProposalBody{
		ChaincodeName:       body.ChaincodeName,
		FunctionName:        body.FunctionName,
		SignedProposalBytes: signedProposal.SignedProposalBytes,
	}

	req.Header = c.GetHeader()
	req.Body = resBody

	res := &ccres.QueryChaincodeResData{}

	err = c.Call(QueryChaincode, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, ErrCallHasError, QueryChaincode)
	}

	return res, nil
}

func (c *FabricClient) PrepareProposal(body ccreq.InvokeChaincodeReqBody, user *userdata.UserData) (*ccres.PrepareProposalResData, error) {
	var err error
	if user == nil {
		user, err = c.LoadUser(body.UserID)
		if err != nil {
			return nil, errors.WithMessagef(err, ErrUserLoadFailed, body.UserID)
		}
	}

	channelId := c.GetAppInfo().GetChannelId()
	fmt.Printf("MspId: %s\n", user.MspId)
	fmt.Printf("ChannelId: %s\n", channelId)

	signedProposal, err := proposal.CreateSignedProposal(channelId, body.ChaincodeName, body.FunctionName, body.Args, user)
	if err != nil {
		return nil, errors.WithMessagef(err, ErrFailedToCreateSignedProposal)
	}
	fmt.Printf("TxId: %s\n", signedProposal.TxId)

	req := &ccreq.PrepareProposalReqData{}
	req.Header = c.GetHeader()
	req.Body = ccreq.PrepareProposalReqBody{
		SignedProposalBytes: signedProposal.SignedProposalBytes,
	}

	res := &ccres.PrepareProposalResData{}
	err = c.Call(PrepareProposal, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, ErrCallHasError, PrepareProposal)
	}

	res.Body.TxID = signedProposal.TxId
	return res, nil
}

func (c *FabricClient) InvokeChaincode(body ccreq.InvokeChaincodeReqBody, user *userdata.UserData) (*ccres.InvokeChaincodeResData, error) {
	var err error
	if user == nil {
		user, err = c.LoadUser(body.UserID)
		if err != nil {
			return nil, errors.WithMessagef(err, ErrUserLoadFailed, body.UserID)
		}
	}

	prepareProposal, err := c.PrepareProposal(body, user)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to prepare proposal")
	}

	fmt.Printf("MspId: %s", user.MspId)
	channelId := c.GetAppInfo().GetChannelId()
	fmt.Printf("MspId: %s", user.MspId)
	fmt.Printf("ChannelId: %s", channelId)

	signedPayloadBytes, err := proposal.SignPayload(prepareProposal.Body.PayloadBytes, user.PrivateKey)
	if err != nil {
		return nil, errors.WithMessagef(err, ErrFailedToCreateSignedProposal)
	}

	req := &ccreq.SubmitProposalReqData{}
	req.Header = c.GetHeader()
	resBody := ccreq.SubmitProposalReqBody{
		TxID:          prepareProposal.Body.TxID,
		ChaincodeName: body.ChaincodeName,
		FunctionName:  body.FunctionName,
		EnvelopeBytes: signedPayloadBytes,
	}
	req.Body = resBody

	res := &ccres.InvokeChaincodeResData{}
	err = c.Call(SubmitProposal, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, ErrCallHasError, SubmitProposal)
	}

	return res, nil
}
