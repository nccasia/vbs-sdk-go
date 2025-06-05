package chaincode

import (
	"encoding/base64"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type InvokeChaincodeReqData struct {
	base.BaseReqModel
	Body InvokeChaincodeReqBody `json:"body"`
}

type SubmitProposalReqData struct {
	base.BaseReqModel
	Body SubmitProposalReqBody `json:"body"`
}

type PrepareProposalReqData struct {
	base.BaseReqModel
	Body PrepareProposalReqBody `json:"body"`
}

type PrepareProposalReqBody struct {
	SignedProposalBytes []byte `json:"signed_proposal_bytes"`
}

type SubmitProposalReqBody struct {
	TxID          string `json:"tx_id"`
	ChaincodeName string `json:"chaincode_name"`
	FunctionName  string `json:"function_name"`
	EnvelopeBytes []byte `json:"envelope_bytes"`
}

type InvokeChaincodeReqBody struct {
	UserID        string   `json:"user_id"`
	ChaincodeName string   `json:"chaincode_name"`
	FunctionName  string   `json:"function_name"`
	Args          []string `json:"args"`
}

func (f *InvokeChaincodeReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.ChaincodeName
	fp = fp + f.Body.FunctionName
	fp = fp + f.Body.UserID

	for _, a := range f.Body.Args {
		fp = fp + a
	}

	return fp
}

func (f *SubmitProposalReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.TxID
	fp = fp + f.Body.ChaincodeName
	fp = fp + f.Body.FunctionName
	fp = fp + base64.StdEncoding.EncodeToString(f.Body.EnvelopeBytes)

	return fp
}

func (f *PrepareProposalReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	fp = fp + base64.StdEncoding.EncodeToString(f.Body.SignedProposalBytes)
	return fp
}
