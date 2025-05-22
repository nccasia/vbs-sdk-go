package event

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type RegisterEventReqData struct {
	base.BaseReqModel
	Body RegisterEventReqBody `json:"body"`
}

type RegisterEventReqBody struct {
	UserID          string `json:"userId"`
	ContractAddress string `json:"contractAddress"`
	EventName       string `json:"eventName"`
	NotifyUrl       string `json:"notifyUrl"`
	AttachArgs      string `json:"attachArgs"`
}

func (f *RegisterEventReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.UserID
	fp = fp + f.Body.ContractAddress
	fp = fp + f.Body.EventName
	fp = fp + f.Body.NotifyUrl
	fp = fp + f.Body.AttachArgs

	return fp
}
