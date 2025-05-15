package base

type ReqInterface interface {
	SetMac(mac string)
	GetEncryptionValue() string
}

type ResInterface interface {
	GetMac() string
	GetEncryptionValue() string
}
