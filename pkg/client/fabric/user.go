package fabric

import (
	userreq "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/user"
	userres "github.com/nccasia/vbs-sdk-go/pkg/core/model/res/fabric/user"
	"github.com/pkg/errors"
)

const (
	RegisterUser = "user/register"
	EnrollUser   = "user/enroll"
)

// RegisterUser register sub user
func (c *FabricClient) RegisterUser(body userreq.UserRegisterReqDataBody) (*userres.UserRegisterResData, error) {
	req := &userreq.UserRegisterReqData{}
	req.Header = c.GetHeader()
	req.Body = body
	res := &userres.UserRegisterResData{}

	err := c.Call(RegisterUser, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", RegisterUser)
	}

	return res, nil
}

// EnrollUser enroll sub user certificate
func (c *FabricClient) EnrollUser(body userreq.UserEnrollReqDataBody) (*userres.UserEnrollResData, error) {
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
