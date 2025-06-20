package keystore

import (
	"os"
	"path"
	"sync"

	"github.com/nccasia/vbs-sdk-go/pkg/common/file"
	"github.com/wonderivan/logger"
)

func NewFileKeyStore(keyDir string) KeyStore {
	file.CheckDir(keyDir)
	return &fileKeyStore{
		path: keyDir,
	}
}

type fileKeyStore struct {
	path string
	m    sync.Mutex
}

func (f *fileKeyStore) keyPath(alias string) string {
	return path.Join(f.path, alias+"_sk")
}

func (f *fileKeyStore) StoreKey(rawPem []byte, alias string) error {
	f.m.Lock()
	defer f.m.Unlock()

	keyFile := f.keyPath(alias)

	err := os.WriteFile(keyFile, rawPem, 0600)
	if err != nil {
		logger.Debug("Failed storing private key [%s]: [%s]", keyFile, err)
		return err
	}
	return nil
}

func (f *fileKeyStore) LoadKey(alias string) ([]byte, error) {
	f.m.Lock()
	defer f.m.Unlock()

	keyFile := f.keyPath(alias)
	raw, err := os.ReadFile(keyFile)
	if err != nil {
		logger.Error("Failed loading private key [%s]: [%s].", alias, err.Error())

		return nil, err
	}
	return raw, nil
}
