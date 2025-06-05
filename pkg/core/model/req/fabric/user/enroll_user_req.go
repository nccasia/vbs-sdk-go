package user

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type UserEnrollReqData struct {
	base.BaseReqModel
	Body UserEnrollReqDataBody `json:"body"`
}

type UserEnrollReqDataBody struct {
	UserID string `json:"userId"`
	Secret string `json:"secret"`
	CSR    string `json:"certificate_request"` // CSR in PEM format
}

func (f *UserEnrollReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.UserID
	fp = fp + f.Body.Secret
	fp = fp + f.Body.CSR

	return fp
}
