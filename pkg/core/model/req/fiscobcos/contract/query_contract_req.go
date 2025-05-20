package chaincode

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type QueryContractReqData struct {
	base.BaseReqModel
	Body QueryContractReqBody `json:"body"`
}

type QueryContractReqBody struct {
	UserID          string   `json:"userId" validate:"required"`
	ContractAddress string   `json:"contractAddress" validate:"required"`
	FunctionName    string   `json:"functionName" validate:"required"`
	Args            []string `json:"args"`
}

func (f *QueryContractReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.UserID
	fp = fp + f.Body.ContractAddress
	fp = fp + f.Body.FunctionName

	for _, a := range f.Body.Args {
		fp = fp + a
	}

	return fp
}
