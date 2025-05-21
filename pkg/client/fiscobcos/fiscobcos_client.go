package fiscobcos

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/nccasia/vbs-sdk-go/pkg/common/encrypt"
	"github.com/nccasia/vbs-sdk-go/pkg/common/http"
	"github.com/nccasia/vbs-sdk-go/pkg/core/config"
	"github.com/nccasia/vbs-sdk-go/pkg/core/constants"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
)

const (
	fiscobcosApiFormat = "%s/api/%s"
)

type FiscoBcosClient struct {
	Config *config.Config
}

func NewFiscoBcosClient(config *config.Config) (*FiscoBcosClient, error) {
	fiscoBcosClient := &FiscoBcosClient{
		Config: config,
	}
	return fiscoBcosClient, nil
}

func (c *FiscoBcosClient) GetHeader() base.ReqHeader {
	return c.Config.GetReqHeader()
}

func (c *FiscoBcosClient) Sign(data string) (string, error) {
	mac, err := encrypt.SignData(constants.Secp256k1, []byte(c.Config.GetAppCert().UserAppPrivateCert), []byte(data))
	if err != nil {
		return "", errors.WithMessage(err, "exception in signature")
	}

	return base64.StdEncoding.EncodeToString(mac), nil
}

func (c *FiscoBcosClient) Call(method string, req base.ReqInterface, res base.ResInterface) error {
	url := c.methodUrl(method)
	mac, err := c.Sign(req.GetEncryptionValue())
	if err != nil {
		return err
	}

	req.SetMac(mac)
	reqBytes, err := json.Marshal(req)
	if err != nil {
		logger.Error("request parameter serialization failed：", err)
		return errors.WithMessage(err, "request parameter serialization failed")
	}

	resBytes, err := http.SendPost(reqBytes, url)
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return errors.WithMessage(err, "send post has error")
	}

	err = json.Unmarshal(resBytes, res)
	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return errors.WithMessage(err, "return parameter serialization failed")
	}

	return nil
}

func (c *FiscoBcosClient) methodUrl(method string) string {
	return fmt.Sprintf(fiscobcosApiFormat, c.Config.GetNodeApi(), method)
}

// func (c *FiscoBcosClient) getBlockLimit() (*big.Int, error) {
// 	res, err := c.GetBlockHeight()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if res.Header.Code != 0 {
// 		return nil, errors.New(res.Header.Msg)
// 	}

// 	height, err := strconv.ParseInt(res.Body.Data, 10, 64)
// 	if err != nil {
// 		return nil, errors.New("ledger height has error")
// 	}

// 	height = height + 100
// 	return new(big.Int).SetInt64(height), nil
// }
