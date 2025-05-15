package res

import (
	"strconv"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type AppInfoResData struct {
	base.BaseResModel
	Body AppInfoResDataBody `json:"body"`
}

type AppInfoResDataBody struct {
	AppType       string `json:"appType"`
	CaType        int    `json:"caType"`
	AlgorithmType int    `json:"algorithmType"`
	MspId         string `json:"mspId"`
	ChannelId     string `json:"channelId"`
	Version       string `json:"version"`
	FabricVersion string `json:"fabricVersion,omitempty"`
}

func (f *AppInfoResData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.AppType
	fp = fp + strconv.Itoa(f.Body.CaType)
	fp = fp + strconv.Itoa(f.Body.AlgorithmType)
	fp = fp + f.Body.MspId
	fp = fp + f.Body.ChannelId
	fp = fp + f.Body.Version
	fp = fp + f.Body.FabricVersion
	return fp
}
