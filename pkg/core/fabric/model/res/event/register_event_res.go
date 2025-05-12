package event

import "github.com/nccasia/vbs-sdk-go/pkg/core/fabric/base"

type RegisterEventResData struct {
	base.BaseResModel
	Body *RegisterEventResBody `json:"body"`
}

type RegisterEventResBody struct {
	EventId string `json:"eventId"`
}

func (f *RegisterEventResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}
	return f.GetBaseEncryptionValue() + f.Body.EventId
}
