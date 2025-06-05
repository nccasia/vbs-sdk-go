package userstore

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/nccasia/vbs-sdk-go/pkg/common/file"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/userdata"
)

func NewUserCertStore(path string) UserCertStore {
	file.CheckDir(path)
	us := &FileUserCertStore{
		FilePath: path,
	}
	return us
}

type FileUserCertStore struct {
	FilePath string
}

func (f *FileUserCertStore) Load(user *userdata.UserData) error {
	key := storeKeyName(user)
	filePath := path.Join(f.FilePath, key)

	if _, err1 := os.Stat(filePath); os.IsNotExist(err1) {
		return errors.New("user not found")
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if bytes == nil {
		return errors.New("user not found")
	}
	user.EnrollmentCertificate = bytes
	return nil
}

func (f *FileUserCertStore) Store(user *userdata.UserData) error {
	key := storeKeyName(user)

	pathFile := path.Join(f.FilePath, key)
	valueBytes := user.EnrollmentCertificate

	err := os.MkdirAll(path.Dir(pathFile), 0700)
	if err != nil {
		return err
	}
	return os.WriteFile(pathFile, valueBytes, 0600)

}

func (f *FileUserCertStore) LoadAll(appCode string) []*userdata.UserData {
	var users []*userdata.UserData

	//Traverse files under the folder
	files, err := os.ReadDir(f.FilePath)
	if err != nil {
		return users
	}

	for _, file := range files {
		filePath := path.Join(f.FilePath, file.Name())

		//get the file name
		name := getPemName(file.Name(), appCode)
		if name != "" {
			//get
			user := &userdata.UserData{}
			bytes, err := os.ReadFile(filePath)
			if err == nil && bytes != nil {
				user.EnrollmentCertificate = bytes
				user.UserName = name
				user.AppCode = appCode
				users = append(users, user)
			}
		}
	}

	return users
}

func getPemName(name, appCode string) string {
	ext := "@" + appCode + "-cert.pem"
	i := strings.Index(name, ext)
	if i != -1 {
		return name[:i]
	} else {
		return ""
	}
}

func storeKeyName(user *userdata.UserData) string {
	return user.UserName + "@" + user.AppCode + "-cert.pem"
}
