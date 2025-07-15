package fiscobcos

import (
	userreq "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fiscobcos/user"
	userres "github.com/nccasia/vbs-sdk-go/pkg/core/model/res/fiscobcos/user"
	"github.com/pkg/errors"
)

const (
	RegisterUser = "user/register"
	EnrollUser   = "user/enroll"
)

func (c *FiscoBcosClient) RegisterUser(body userreq.UserRegisterReqDataBody) (*userres.UserRegisterResData, error) {
	req := &userreq.UserRegisterReqData{}
	req.Header = c.GetHeader()
	req.Body = body
	res := &userres.UserRegisterResData{}

	err := c.Call(RegisterUser, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *FiscoBcosClient) EnrollUser(body userreq.UserEnrollReqDataBody) (*userres.UserEnrollResData, error) {
	req := &userreq.UserEnrollReqData{}
	req.Header = c.GetHeader()
	req.Body = body
	res := &userres.UserEnrollResData{}

	err := c.Call(EnrollUser, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", EnrollUser)
	}

	return res, nil
}
