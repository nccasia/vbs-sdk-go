package user

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type UserEnrollResData struct {
	base.BaseResModel
	Body *UserEnrollResDataBody `json:"body"`
}

type UserEnrollResDataBody struct {
	Cert string `json:"cert"`
}

func (f *UserEnrollResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	return f.GetBaseEncryptionValue() + f.Body.Cert
}
