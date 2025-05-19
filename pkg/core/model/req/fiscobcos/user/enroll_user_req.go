package user

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type UserEnrollReqData struct {
	base.BaseReqModel
	Body UserEnrollReqDataBody `json:"body"`
}

type UserEnrollReqDataBody struct {
	UserID string `json:"userId"`
}

func (f *UserEnrollReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.UserID
	return fp
}
