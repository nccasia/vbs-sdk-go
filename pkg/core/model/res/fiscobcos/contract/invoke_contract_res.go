package chaincode

import (
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type InvokeContractResData struct {
	base.BaseResModel
	Body *InvokeContractResBody `json:"body"`
}

type InvokeContractResBody struct {
	Status          string `json:"status"`
	TransactionHash string `json:"transactionHash"`
}

func (f *InvokeContractResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.Status
	fp = fp + f.Body.TransactionHash
	return fp
}
