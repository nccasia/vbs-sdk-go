package fabric

import (
	"github.com/nccasia/vbs-sdk-go/pkg/client/fabric/node"
	"github.com/nccasia/vbs-sdk-go/pkg/common/keystore"
	"github.com/nccasia/vbs-sdk-go/pkg/common/userstore"
	"github.com/nccasia/vbs-sdk-go/pkg/core/config"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/userdata"

	"github.com/wonderivan/logger"
)

type FabricClient struct {
	appInfo         config.AppInfo
	userCode        string
	nodeClients     map[string]*node.NodeClient
	defaultNodeName string
	keyOpts         keystore.KeyStore
	userOpts        userstore.UserCertStore

	users map[string]*userdata.UserData
}

func InitFabricClient(config *config.Config) (*FabricClient, error) {
	if err := config.Init(); err != nil {
		logger.Error("configuration initialization failed")
		return nil, err
	}

	defCli := node.NewNodeCli(config.GetNodeApi(), config.GetAppCert().UserAppPrivateCert)

	fabricClient := &FabricClient{
		appInfo:     config.GetAppInfo(),
		userCode:    config.GetUserCode(),
		nodeClients: make(map[string]*node.NodeClient),
		users:       make(map[string]*userdata.UserData),
	}

	if fabricClient.keyOpts == nil {
		fabricClient.keyOpts = keystore.NewFileKeyStore(config.GetKSPath())
	}

	if fabricClient.userOpts == nil {
		fabricClient.userOpts = userstore.NewUserCertStore(config.GetUSPath())
	}

	if fabricClient.defaultNodeName == "" {
		fabricClient.defaultNodeName = fabricClient.appInfo.MspId
	}

	fabricClient.nodeClients[fabricClient.defaultNodeName] = defCli
	return fabricClient, nil
}

func (c *FabricClient) GetHeader() base.ReqHeader {
	header := base.ReqHeader{
		UserCode: c.userCode,
		AppCode:  c.appInfo.AppCode,
	}

	return header
}

func (c *FabricClient) GetAppInfo() *config.AppInfo {
	return &c.appInfo
}

func (c *FabricClient) newUser(userName string) *userdata.UserData {
	user := &userdata.UserData{
		UserName: userName,
		AppCode:  c.appInfo.AppCode,
		MspId:    c.appInfo.MspId,
	}

	return user
}

func (c *FabricClient) Call(method string, req base.ReqInterface, res base.ResInterface) error {
	return c.nodeClients[c.defaultNodeName].Call(method, req, res)
}
