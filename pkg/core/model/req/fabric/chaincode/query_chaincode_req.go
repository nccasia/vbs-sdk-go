package chaincode

import (
	"encoding/base64"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type QueryChaincodeReqData struct {
	base.BaseReqModel
	Body SignedProposalBody `json:"body"`
}

type SignedProposalBody struct {
	ChaincodeName       string `json:"chaincode_name"`
	FunctionName        string `json:"function_name"`
	SignedProposalBytes []byte `json:"signed_proposal_bytes"`
}
type QueryChaincodeReq struct {
	ChaincodeName string   `json:"chaincode_name"`
	FunctionName  string   `json:"function_name"`
	UserID        string   `json:"user_id"`
	Args          []string `json:"args"`
}

func (f *QueryChaincodeReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.ChaincodeName
	fp = fp + f.Body.FunctionName
	fp = fp + base64.StdEncoding.EncodeToString(f.Body.SignedProposalBytes)

	return fp
}
