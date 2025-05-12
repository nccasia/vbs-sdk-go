package event

import "github.com/nccasia/vbs-sdk-go/pkg/core/fabric/base"

type QueryEventResData struct {
	base.BaseResModel
	Body []QueryEventResBody `json:"body"`
}

type QueryEventResBody struct {
	EventId    string `json:"eventId"`
	EventKey   string `json:"eventKey,omitempty"`
	NotifyUrl  string `json:"notifyUrl"`
	AttachArgs string `json:"attachArgs"`
	CreateTime string `json:"createTime"`
	OrgCode    string `json:"orgCode"`
	UserCode   string `json:"userCode"`
	AppCode    string `json:"appCode"`
	ChainCode  string `json:"chainCode,omitempty"`
	EventType  string `json:"eventType,omitempty"`
	UserID     string `json:"userId"`
}

func (f *QueryEventResData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	for _, task := range f.Body {
		fp = fp + task.EventId
		fp = fp + task.EventKey
		fp = fp + task.NotifyUrl
		fp = fp + task.AttachArgs
		fp = fp + task.CreateTime
		fp = fp + task.OrgCode
		fp = fp + task.UserCode
		fp = fp + task.AppCode
		fp = fp + task.ChainCode
		fp = fp + task.EventType
		fp = fp + task.UserID
	}

	return fp
}
