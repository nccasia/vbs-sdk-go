package chaincode

import (
	"encoding/base64"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type QueryContractResData struct {
	base.BaseResModel
	Body *QueryContractResBody `json:"body"`
}

type QueryContractResBody struct {
	Status  string `json:"status"`
	Payload []byte `json:"payload"`
}

func (f *QueryContractResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.Status
	fp = fp + base64.StdEncoding.EncodeToString(f.Body.Payload)
	return fp
}
