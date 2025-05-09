package chaincode

import "vbs-sdk-go/pkg/core/fabric/base"

type QueryChaincodeReqData struct {
	base.BaseReqModel
	Body QueryChaincodeReqBody `json:"body"`
}

type QueryChaincodeReqBody struct {
	ChaincodeName string   `json:"chaincode_name"`
	FunctionName  string   `json:"function_name"`
	UserID        string   `json:"user_id"`
	Args          []string `json:"args"`
}

func (f *QueryChaincodeReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.ChaincodeName
	fp = fp + f.Body.FunctionName
	fp = fp + f.Body.UserID

	for _, a := range f.Body.Args {
		fp = fp + a
	}

	return fp
}
