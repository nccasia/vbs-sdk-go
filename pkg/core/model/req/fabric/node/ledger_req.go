package node

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type LedgerReqData struct {
	base.BaseReqModel
	Body LedgerReqDataBody `json:"body"`
}

type LedgerReqDataBody struct {
}

func (f *LedgerReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue()
}
