package event

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type QueryEventResData struct {
	base.BaseResModel
	Body []QueryEventResBody `json:"body"`
}

type QueryEventResBody struct {
	EventId         string `json:"eventId"`
	EventName       string `json:"eventName,omitempty"`
	AppCode         string `json:"appCode"`
	UserCode        string `json:"userCode"`
	NotifyUrl       string `json:"notifyUrl"`
	AttachArgs      string `json:"attachArgs"`
	CreateTime      string `json:"createTime"`
	ContractAddress string `json:"contractAddress,omitempty"`
	EventType       string `json:"eventType"`
	UserID          string `json:"userId"`
}

func (f *QueryEventResData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	for _, task := range f.Body {
		fp = fp + task.EventId
		fp = fp + task.EventName
		fp = fp + task.AppCode
		fp = fp + task.UserCode
		fp = fp + task.NotifyUrl
		fp = fp + task.AttachArgs
		fp = fp + task.CreateTime
		fp = fp + task.ContractAddress
		fp = fp + task.EventType
		fp = fp + task.UserID
	}

	return fp
}
