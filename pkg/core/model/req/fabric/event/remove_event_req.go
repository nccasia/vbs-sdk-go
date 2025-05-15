package event

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type RemoveEventReqData struct {
	base.BaseReqModel
	Body RemoveEventReqBody `json:"body"`
}

type RemoveEventReqBody struct {
	EventId string `json:"eventId"`
	UserID  string `json:"userId"`
}

func (f *RemoveEventReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.EventId + f.Body.UserID
}
