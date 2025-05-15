package user

import (
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type UserRegisterReqData struct {
	base.BaseReqModel
	Body UserRegisterReqDataBody `json:"body"`
}

type UserRegisterReqDataBody struct {
	UserID      string `json:"userId"`
	Affiliation string `json:"affiliation,optional"`
	Attributes  string `json:"attributes,optional"`
}

func (f *UserRegisterReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.UserID
	fp = fp + f.Body.Affiliation
	fp = fp + f.Body.Attributes
	return fp
}
