package userstore

import "github.com/nccasia/vbs-sdk-go/pkg/core/model/userdata"

type UserCertStore interface {
	Load(user *userdata.UserData) error
	LoadAll(appCode string) []*userdata.UserData
	Store(user *userdata.UserData) error
}
