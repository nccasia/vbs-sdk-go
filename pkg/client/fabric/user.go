package fabric

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/cloudflare/cfssl/csr"
	"github.com/nccasia/vbs-sdk-go/pkg/common/encrypt"
	"github.com/nccasia/vbs-sdk-go/pkg/core/constants"
	userreq "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fabric/user"
	userres "github.com/nccasia/vbs-sdk-go/pkg/core/model/res/fabric/user"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/userdata"
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

	privKey, err := encrypt.GeneratePrivateKey(constants.Prime256v1)
	if err != nil {
		return nil, errors.WithMessage(err, "Generate private key has error")
	}

	csrReq := encrypt.NewCertificateRequest(body.UserID)
	csrBytes, err := csr.Generate(privKey, csrReq)
	if err != nil {
		return nil, errors.WithMessage(err, "Generate csr has error")
	}

	req.Body.CSR = string(csrBytes)

	// save key
	rawKey, err := encrypt.PrivateKeyToPEM(privKey)
	if err != nil {
		return nil, errors.WithMessage(err, "privateKey to PEM has error")
	}
	err = c.keyOpts.StoreKey(rawKey, hex.EncodeToString(encrypt.PivSKI(privKey)))
	if err != nil {
		return nil, errors.WithMessage(err, "save privateKey has error")
	}

	res := &userres.UserEnrollResData{}

	err = c.Call(EnrollUser, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", EnrollUser)
	}

	user := c.newUser(body.UserID)
	user.PrivateKey = privKey
	certBytes, err := base64.StdEncoding.DecodeString(string(res.Body.Cert))
	if err != nil {
		return nil, errors.WithMessagef(err, "decode user cert error")
	}
	user.EnrollmentCertificate = certBytes
	err = c.userOpts.Store(user)
	if err != nil {
		return nil, errors.WithMessagef(err, "store user cert error")
	}

	c.users[body.UserID] = user

	return res, nil
}

// LoadUser load user from local store , before, the cache is checked from the client users
func (c *FabricClient) LoadUser(userId string) (*userdata.UserData, error) {
	if userId == "" {
		return nil, errors.New("userName can not be empty")
	}
	user, ok := c.users[userId]
	if ok {
		return user, nil
	}

	user = c.newUser(userId)
	err := c.userOpts.Load(user)
	if err != nil {
		return nil, err
	}

	puk, err := encrypt.LoadPublicKeyByCertPem(string(user.EnrollmentCertificate))
	if err != nil {
		return nil, errors.WithMessage(err, "cert pem load has error")
	}

	alias := hex.EncodeToString(encrypt.PubSKI(puk))
	fmt.Printf("alias: %s \n", alias)

	bytesKey, err := c.keyOpts.LoadKey(alias)
	if err != nil {
		return nil, errors.WithMessage(err, "load private key has error")
	}

	k, err := encrypt.LoadPrivateKeyFromPEM(bytesKey)
	if err != nil {
		return nil, errors.WithMessage(err, "new private key provider has error")
	}

	user.PrivateKey = k
	c.users[userId] = user
	return user, nil
}
