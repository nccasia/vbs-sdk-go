package chaincode

import (
	"encoding/base64"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type InvokeChaincodeResData struct {
	base.BaseResModel
	Body *InvokeChaincodeResBody `json:"body"`
}

type PrepareProposalResData struct {
	base.BaseResModel
	Body *PrepareProposalResBody `json:"body"`
}

type PrepareProposalResBody struct {
	TxID         string `json:"tx_id"`
	PayloadBytes []byte `json:"payload_bytes"`
}

type InvokeChaincodeResBody struct {
	TxID    string `json:"tx_id"`
	Status  string `json:"status"`
	Payload []byte `json:"payload"`
}

func (f *InvokeChaincodeResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.TxID
	fp = fp + f.Body.Status
	fp = fp + base64.StdEncoding.EncodeToString(f.Body.Payload)
	return fp
}

func (f *PrepareProposalResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()

	fp = fp + base64.StdEncoding.EncodeToString(f.Body.PayloadBytes)
	return fp
}
