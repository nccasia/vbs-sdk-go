package event

import "github.com/nccasia/vbs-sdk-go/pkg/core/fabric/base"

type QueryEventReqData struct {
	base.BaseReqModel
	Body QueryEventReqBody `json:"body"`
}

type QueryEventReqBody struct {
	UserID string `json:"userId"`
}

func (f *QueryEventReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.UserID

	return fp
}
