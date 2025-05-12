package event

import "github.com/nccasia/vbs-sdk-go/pkg/core/fabric/base"

type RegisterEventReqData struct {
	base.BaseReqModel
	Body RegisterEventReqBody `json:"body"`
}

type RegisterEventReqBody struct {
	UserID     string `json:"userId"`
	ChainCode  string `json:"chainCode"`
	EventKey   string `json:"eventKey"`
	NotifyUrl  string `json:"notifyUrl"`
	AttachArgs string `json:"attachArgs"`
}

func (f *RegisterEventReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.UserID
	fp = fp + f.Body.ChainCode
	fp = fp + f.Body.EventKey
	fp = fp + f.Body.NotifyUrl
	fp = fp + f.Body.AttachArgs

	return fp
}
