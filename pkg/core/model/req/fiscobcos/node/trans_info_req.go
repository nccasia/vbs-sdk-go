package node

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type TransInfoReqData struct {
	base.BaseReqModel
	Body TransInfoReqDataBody `json:"body"`
}

type TransInfoReqDataBody struct {
	TxHash string `json:"txHash"`
}

func (f *TransInfoReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.TxHash
}
