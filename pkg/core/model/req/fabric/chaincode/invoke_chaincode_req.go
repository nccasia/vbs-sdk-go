package chaincode

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type InvokeChaincodeReqData struct {
	base.BaseReqModel
	Body InvokeChaincodeReqBody `json:"body"`
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
