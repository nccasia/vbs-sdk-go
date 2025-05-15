package req

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type AppInfoReqData struct {
	base.BaseReqModel
}

func (f *AppInfoReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue()
}
