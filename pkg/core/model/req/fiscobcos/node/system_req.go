package node

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type SystemReqData struct {
	base.BaseReqModel
}

func (f *SystemReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue()
}
