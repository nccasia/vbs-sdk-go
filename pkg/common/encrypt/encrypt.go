package encrypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nccasia/vbs-sdk-go/pkg/core/constants"
)

type pkcs8Info struct {
	Version             int
	PrivateKeyAlgorithm []asn1.ObjectIdentifier
	PrivateKey          []byte
}

type ecPrivateKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

var (
	// curveHalfOrders contains the precomputed curve group orders halved.
	// It is used to ensure that signature' S value is lower or equal to the
	// curve group order halved. We accept only low-S signatures.
	// They are precomputed for efficiency reasons.
	curveHalfOrders = map[elliptic.Curve]*big.Int{
		elliptic.P224(): new(big.Int).Rsh(elliptic.P224().Params().N, 1),
		elliptic.P256(): new(big.Int).Rsh(elliptic.P256().Params().N, 1),
		elliptic.P384(): new(big.Int).Rsh(elliptic.P384().Params().N, 1),
		elliptic.P521(): new(big.Int).Rsh(elliptic.P521().Params().N, 1),
		crypto.S256():   new(big.Int).Rsh(crypto.S256().Params().N, 1),
	}
)

func SignData(encryptType string, priKeyPem, digest []byte) ([]byte, error) {
	priKey, err := LoadPrivateKeyFromPEM(encryptType, priKeyPem)
	if err != nil {
		return nil, errors.New("could not parse private key: " + err.Error())
	}

	// Hash data by SHA-256
	hash := sha256.Sum256(digest)
	return SignECDSA(priKey, hash[:])
}

// LoadPrivateKeyFromPEM loads an ECDSA private key from PEM-encoded data.
func LoadPrivateKeyFromPEM(encryptType string, pemData []byte) (*ecdsa.PrivateKey, error) {
	// Decode the PEM block
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// Check the block type
	if block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("invalid PEM block type: %s, expected PRIVATE KEY", block.Type)
	}

	if encryptType == constants.Secp256k1 {
		var pkcs8Key pkcs8Info
		if _, err := asn1.Unmarshal(block.Bytes, &pkcs8Key); err != nil {
			return nil, err
		}
		var privateKey ecPrivateKey
		if _, err := asn1.Unmarshal(pkcs8Key.PrivateKey, &privateKey); err != nil {
			return nil, err
		}
		return crypto.ToECDSA(privateKey.PrivateKey)
	} else {
		// Parse the private key from DER-encoded data
		priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)

		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %v", err)
		}

		// Ensure the parsed key is an ECDSA private key
		ecdsaPriv, ok := priv.(*ecdsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("private key is not an ECDSA key")
		}

		return ecdsaPriv, nil
	}
}

func SignECDSA(k *ecdsa.PrivateKey, digest []byte) (signature []byte, err error) {
	r, s, err := ecdsa.Sign(rand.Reader, k, digest)
	if err != nil {
		return nil, err
	}

	s, _, err = ToLowS(&k.PublicKey, s)
	if err != nil {
		return nil, err
	}

	return marshalECDSASignature(r, s)
}

func ToLowS(k *ecdsa.PublicKey, s *big.Int) (*big.Int, bool, error) {
	lowS, err := IsLowS(k, s)
	if err != nil {
		return nil, false, err
	}

	if !lowS {
		// Set s to N - s that will be then in the lower part of signature space
		// less or equal to half order
		s.Sub(k.Params().N, s)

		return s, true, nil
	}

	return s, false, nil
}

func IsLowS(k *ecdsa.PublicKey, s *big.Int) (bool, error) {
	halfOrder, ok := curveHalfOrders[k.Curve]
	if !ok {
		return false, fmt.Errorf("curve not recognized [%s]", k.Curve)
	}

	return s.Cmp(halfOrder) != 1, nil
}

func marshalECDSASignature(r, s *big.Int) ([]byte, error) {
	return asn1.Marshal(ECDSASignature{r, s})
}
