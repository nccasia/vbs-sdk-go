package event

import "vbs-sdk-go/pkg/core/fabric/base"

type RemoveEventResData struct {
	base.BaseResModel
}

func (f *RemoveEventResData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue()
}
