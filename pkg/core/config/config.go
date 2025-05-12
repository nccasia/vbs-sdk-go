package config

import (
	"github.com/nccasia/vbs-sdk-go/pkg/core/fabric/base"
)

type CertInfo struct {
	//public key cert of DApp
	AppPublicCert string

	//Private key cert of user
	UserAppPrivateCert string
}

type Config struct {
	nodeApi string
	mspDir  string
	user    userInfo
	app     AppInfo

	appCert CertInfo
	isInit  bool
}

type AppInfo struct {
	AppCode string
	AppType string

	// CAType        enum.App_CaType
	// AlgorithmType enum.App_AlgorithmType

	MspId     string
	ChannelId string
	Version   string
}

type userInfo struct {
	UserCode string
}

func NewConfig(api, userCode, appCode, prk, mspDir string) (*Config, error) {
	config := &Config{
		nodeApi: api,
		mspDir:  mspDir,
		appCert: CertInfo{AppPublicCert: "", UserAppPrivateCert: prk},
		user:    userInfo{UserCode: userCode},
		app:     AppInfo{AppCode: appCode},
	}
	err := config.Init()
	return config, err
}

func (c *Config) Init() error {
	if !c.isInit {
		// reqData := apptypes.AppInfoReqData{}
		// reqData.Header = c.GetReqHeader()

		// res, err := app.GetAppInfo(&reqData, c.nodeApi, "")

		// if err != nil {
		// 	return err
		// }

		// if res.Header.Code != 0 {
		// 	return errors.New("get app info failed ：" + res.Header.Msg)
		// }

		// c.app.AppType = res.Body.AppType

		// c.app.CAType = enum.App_CaType(res.Body.CaType)
		// c.app.AlgorithmType = enum.App_AlgorithmType(res.Body.AlgorithmType)

		// if c.appCert.AppPublicCert == "" {
		// 	c.appCert.AppPublicCert = GetGatewayPublicKey(c.app.AlgorithmType)
		// }

		// if c.appCert.AppPublicCert == "" {
		// 	return errors.New("gateway public key not setting")
		// }

		c.app.MspId = "Org1MSP"
		// c.app.MspId = res.Body.MspId
		// c.app.ChannelId = res.Body.ChannelId
		// c.app.Version = res.Body.FabricVersion
		c.isInit = true
	}

	return nil
}

func (c *Config) GetReqHeader() base.ReqHeader {
	header := base.ReqHeader{
		UserCode: c.user.UserCode,
		AppCode:  c.app.AppCode,
	}

	return header
}

func (c *Config) GetNodeApi() string {
	return c.nodeApi
}

func (c *Config) GetAppInfo() AppInfo {
	return c.app
}

func (c *Config) GetAppCert() CertInfo {
	return c.appCert
}

func (c *Config) GetUserCode() string {
	return c.user.UserCode
}
