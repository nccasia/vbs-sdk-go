package event

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

type RemoveEventResData struct {
	base.BaseResModel
}

func (f *RemoveEventResData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue()
}
