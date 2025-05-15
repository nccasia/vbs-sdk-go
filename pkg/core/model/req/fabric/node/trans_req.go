package node

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type TransReqData struct {
	base.BaseReqModel
	Body TransReqDataBody `json:"body"`
}

type TransReqDataBody struct {
	TxId     string `json:"txId"`
	DataType string `json:"dataType,optional"`
}

func (f *TransReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.TxId + f.Body.DataType
}
