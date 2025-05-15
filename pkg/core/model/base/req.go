package base

type ReqHeader struct {
	UserCode string `json:"userCode"` // User unique identification
	AppCode  string `json:"appCode"`  // App unique identification
}

type BaseReqModel struct {
	Header ReqHeader `json:"header"`
	Mac    string    `json:"mac"`
}

func (b *BaseReqModel) SetMac(mac string) {
	b.Mac = mac
}

func (b *BaseReqModel) GetBaseEncryptionValue() string {
	return b.Header.UserCode + b.Header.AppCode
}
