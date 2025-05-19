package user

import (
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type UserRegisterReqData struct {
	base.BaseReqModel
	Body UserRegisterReqDataBody `json:"body"`
}

type UserRegisterReqDataBody struct {
	UserID string `json:"userId"`
}

func (f *UserRegisterReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.UserID
	return fp
}
