package file

import (
	"os"

	"github.com/wonderivan/logger"
)

func CheckDir(path string) {
	if _, err1 := os.Stat(path); !os.IsNotExist(err1) {
		return
	}

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		logger.Warn("path [%s] create failed : %s", path, err.Error())
	}
}
