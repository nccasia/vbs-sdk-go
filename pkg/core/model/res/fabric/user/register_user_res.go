package user

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type UserRegisterResData struct {
	base.BaseResModel
	Body *UserRegisterResBody `json:"body"`
}

type UserRegisterResBody struct {
	UserId string `json:"userId"`
	Secret string `json:"secret"`
}

func (f *UserRegisterResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}
	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.UserId
	fp = fp + f.Body.Secret
	return fp
}
