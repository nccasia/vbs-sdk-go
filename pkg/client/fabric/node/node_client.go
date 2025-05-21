package node

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/nccasia/vbs-sdk-go/pkg/common/encrypt"
	"github.com/nccasia/vbs-sdk-go/pkg/common/http"
	"github.com/nccasia/vbs-sdk-go/pkg/core/constants"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"

	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
)

type NodeClient struct {
	gateWayURL string
	// sign       sign.Crypto
	// mspId string
	pri string
}

const (
	fabricApiFormat = "%s/api/%s"
)

func NewNodeCli(url string, pri string) *NodeClient {
	cli := &NodeClient{
		gateWayURL: url,
		pri:        pri,
	}
	return cli
}

func (c *NodeClient) Sign(data string) (string, error) {
	mac, err := encrypt.SignData(constants.Prime256v1, []byte(c.pri), []byte(data))
	if err != nil {
		return "", errors.WithMessage(err, "exception in signature")
	}

	return base64.StdEncoding.EncodeToString(mac), nil
}

func (c *NodeClient) Call(method string, req base.ReqInterface, res base.ResInterface) error {
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

func (c *NodeClient) methodUrl(method string) string {
	return fmt.Sprintf(fabricApiFormat, c.gateWayURL, method)
}
