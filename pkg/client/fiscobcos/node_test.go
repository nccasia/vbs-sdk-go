package fiscobcos

import (
	"testing"

	"github.com/nccasia/vbs-sdk-go/pkg/core/config"
)

func getFiscoBcosClient(t *testing.T) *FiscoBcosClient {
	api := "http://localhost:8888"
	userCode := "FiscoBcosUserCode1"
	appCode := "FiscoBcosAppCode1"
	privK := "-----BEGIN PRIVATE KEY-----\nMIGNAgEAMBAGByqGSM49AgEGBSuBBAAKBHYwdAIBAQQg+sSm14WYvHmyr2Boh2nG\nIkcRnDZlsYsv5D4szeoDT+GgBwYFK4EEAAqhRANCAASVHE4ihBmSYo9alPymFHCT\n/MAvnbmqEFS2vNl/1p/n+6CDtTcanLvTcj0fYwNQNl6KlFJtPTJlym16iax7f8Ww\n-----END PRIVATE KEY-----\n"
	mspDir := "C:\\test"

	config, err := config.NewConfig(api, userCode, appCode, privK, mspDir)
	if err != nil {
		t.Fatal(err)
	}

	fiscobcosClient, err := NewFiscoBcosClient(config)

	if err != nil {
		t.Fatal(err)
	}

	return fiscobcosClient
}
