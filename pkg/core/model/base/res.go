package base

import "strconv"

type ResHeader struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BaseResModel struct {
	Header *ResHeader `json:"header"`
	Mac    string     `json:"mac"`
}

func (b *BaseResModel) GetMac() string {
	return b.Mac
}

func (b *BaseResModel) GetEncryptionValue() string {
	return b.GetBaseEncryptionValue()
}
func (b *BaseResModel) GetBaseEncryptionValue() string {
	return strconv.Itoa(b.Header.Code) + b.Header.Msg
}
