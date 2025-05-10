package chaincode

import (
	"encoding/base64"
	"vbs-sdk-go/pkg/core/fabric/base"
)

type QueryChaincodeResData struct {
	base.BaseResModel
	Body *QueryChaincodeResBody `json:"body"`
}

type QueryChaincodeResBody struct {
	Status  string `json:"status"`
	Payload []byte `json:"payload"`
}

func (f *QueryChaincodeResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.Status
	fp = fp + base64.StdEncoding.EncodeToString(f.Body.Payload)
	return fp
}
