package encrypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"

	"github.com/cloudflare/cfssl/csr"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nccasia/vbs-sdk-go/pkg/core/constants"
)

var (
	oidNamedCurveP256 = asn1.ObjectIdentifier{1, 2, 840, 10045, 3, 1, 7}
	oidNamedCurveS256 = asn1.ObjectIdentifier{1, 3, 132, 0, 10}
)

var oidPublicKeyECDSA = asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}

func oidFromNamedCurve(curve elliptic.Curve) (asn1.ObjectIdentifier, bool) {
	fmt.Printf("CURVE %d", curve)
	switch curve {
	case elliptic.P256():
		return oidNamedCurveP256, true
	case crypto.S256():
		return oidNamedCurveS256, true
	}
	return oidNamedCurveS256, true
}

func namedCurveFromOID(oid asn1.ObjectIdentifier) elliptic.Curve {
	switch {
	case oid.Equal(oidNamedCurveP256):
		return elliptic.P256()
	case oid.Equal(oidNamedCurveS256):
		return crypto.S256()
	}
	return nil
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

func SignData(priKeyPem, digest []byte) ([]byte, error) {
	priKey, err := LoadPrivateKeyFromPEM(priKeyPem)
	if err != nil {
		return nil, errors.New("could not parse private key: " + err.Error())
	}

	// Hash data by SHA-256
	hash := sha256.Sum256(digest)
	return SignECDSA(priKey, hash[:])
}

// LoadPrivateKeyFromPEM loads an ECDSA private key from PEM-encoded data.
func LoadPrivateKeyFromPEM(pemData []byte) (*ecdsa.PrivateKey, error) {
	encryptType, err := DetectEncryptTypeFromPEM(pemData)
	if err != nil {
		return nil, fmt.Errorf("failed to detect encrypt type PEM block")
	}
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

func LoadPublicKeyFromPEM(encryptType string, pemData []byte) (*ecdsa.PublicKey, error) {
	// Decode the PEM block
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// Check the block type
	if block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("invalid PEM block type: %s, expected PUBLIC KEY", block.Type)
	}

	if encryptType == constants.Secp256k1 {
		var pkixPublicKey pkixPublicKey
		if _, err := asn1.Unmarshal(block.Bytes, &pkixPublicKey); err != nil {
			return nil, err
		}

		return crypto.UnmarshalPubkey(pkixPublicKey.BitString.Bytes)
	} else {
		// Parse the public key from DER-encoded data
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse public key: %v", err)
		}

		// Ensure the parsed key is an ECDSA public key
		ecdsaPub, ok := pub.(*ecdsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("public key is not an ECDSA key")
		}

		return ecdsaPub, nil
	}
}

func LoadPublicKeyByCertPem(cert string) (*ecdsa.PublicKey, error) {

	bl, _ := pem.Decode([]byte(cert))
	if bl == nil {
		return nil, errors.New("failed to decode PEM block from Certificate")
	}

	key, err := ParsePublicKeyByCert(bl.Bytes)

	if err != nil {
		return nil, errors.New("failed to parse private key from PrivateKey")
	}

	return key.(*ecdsa.PublicKey), nil
}

func ParsePublicKeyByCert(certBytes []byte) (pub interface{}, err error) {
	var cert certificate
	rest, err := asn1.Unmarshal(certBytes, &cert)
	if err != nil {
		return nil, err
	}
	if len(rest) > 0 {
		return nil, asn1.SyntaxError{Msg: "trailing data"}
	}

	publicKey, err := parsePublicKey(&cert.TBSCertificate.PublicKey)

	return publicKey, err
}

func parsePublicKey(keyData *publicKeyInfo) (interface{}, error) {
	asn1Data := keyData.PublicKey.RightAlign()
	paramsData := keyData.Algorithm.Parameters.FullBytes
	namedCurveOID := new(asn1.ObjectIdentifier)
	rest, err := asn1.Unmarshal(paramsData, namedCurveOID)
	if err != nil {
		return nil, err
	}
	if len(rest) != 0 {
		return nil, errors.New("x509: trailing data after ECDSA parameters")
	}
	namedCurve := namedCurveFromOID(*namedCurveOID)
	if namedCurve == nil {
		return nil, errors.New("x509: unsupported elliptic curve")
	}
	x, y := elliptic.Unmarshal(namedCurve, asn1Data)
	if x == nil {
		return nil, errors.New("x509: failed to unmarshal elliptic curve point")
	}
	pub := &ecdsa.PublicKey{
		Curve: namedCurve,
		X:     x,
		Y:     y,
	}
	return pub, nil
}

func DetectEncryptTypeFromPEM(pemData []byte) (string, error) {
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("invalid PEM block")
	}

	var pkcs8 pkcs8Info
	_, err := asn1.Unmarshal(block.Bytes, &pkcs8)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal PKCS#8: %v", err)
	}

	// The second OID in PrivateKeyAlgorithm is the named curve
	if len(pkcs8.PrivateKeyAlgorithm) < 2 {
		return "", errors.New("missing curve OID in key")
	}

	oid := pkcs8.PrivateKeyAlgorithm[1].String()
	fmt.Println("Detected OID: ", oid)
	switch oid {
	case oidNamedCurveP256.String():
		return constants.Prime256v1, nil
	case oidNamedCurveS256.String():
		return constants.Secp256k1, nil
	default:
		return "unknown (" + oid + ")", nil
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

// PrivateKeyToPEM converts the private key to PEM format.
func PrivateKeyToPEM(k *ecdsa.PrivateKey) ([]byte, error) {
	if k == nil {
		return nil, errors.New("invalid ecdsa private key: nil")
	}

	// Get OID for the curve
	oidNamedCurve, ok := oidFromNamedCurve(k.Curve)
	if !ok {
		return nil, errors.New("unknown elliptic curve")
	}

	// Pad private key
	privateKeyBytes := k.D.Bytes()
	paddedPrivateKey := make([]byte, (k.Curve.Params().N.BitLen()+7)/8)
	copy(paddedPrivateKey[len(paddedPrivateKey)-len(privateKeyBytes):], privateKeyBytes)

	// Serialize public key
	var pubKeyBytes []byte
	if k.Curve == elliptic.P256() {
		// Use ECDH for P-256
		priDchKey, err := k.ECDH()
		if err != nil {
			return nil, fmt.Errorf("failed to convert to ECDH key: %v", err)
		}
		pubKeyBytes = priDchKey.PublicKey().Bytes()
	} else {
		// Manual serialization for other curves (e.g., secp256k1)
		pubKeyBytes = make([]byte, 65)
		pubKeyBytes[0] = 0x04
		xBytes := k.PublicKey.X.Bytes()
		yBytes := k.PublicKey.Y.Bytes()
		copy(pubKeyBytes[1:33], xBytes[len(xBytes)-32:])
		copy(pubKeyBytes[33:], yBytes[len(yBytes)-32:])
	}

	// Marshal to ASN.1 ECPrivateKey structure
	asn1Bytes, err := asn1.Marshal(ecPrivateKey{
		Version:       1,
		PrivateKey:    paddedPrivateKey,
		NamedCurveOID: oidNamedCurve,
		PublicKey:     asn1.BitString{Bytes: pubKeyBytes},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal EC key to ASN.1: %v", err)
	}

	// Create PKCS#8 structure
	pkcs8Key := pkcs8Info{
		Version: 0,
		PrivateKeyAlgorithm: []asn1.ObjectIdentifier{
			oidPublicKeyECDSA,
			oidNamedCurve,
		},
		PrivateKey: asn1Bytes,
	}

	// Marshal to PKCS#8
	pkcs8Bytes, err := asn1.Marshal(pkcs8Key)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal PKCS#8 key to ASN.1: %v", err)
	}

	// Encode to PEM
	return pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pkcs8Bytes,
	}), nil
}

func MarshalPKIXPublicKey(pub *ecdsa.PublicKey) ([]byte, error) {
	var publicKeyBytes []byte
	var publicKeyAlgorithm pkix.AlgorithmIdentifier
	var err error

	if publicKeyBytes, publicKeyAlgorithm, err = marshalPublicKey(pub); err != nil {
		return nil, err
	}

	pukix := pkixPublicKey{
		Algo: publicKeyAlgorithm,
		BitString: asn1.BitString{
			Bytes:     publicKeyBytes,
			BitLength: 8 * len(publicKeyBytes),
		},
	}

	ret, _ := asn1.Marshal(pukix)
	return ret, nil
}

// PublicKeyToPEM marshals a public key to the pem format
func PublicKeyToPEM(k *ecdsa.PublicKey) ([]byte, error) {
	PubASN1, err := MarshalPKIXPublicKey(k)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: PubASN1,
		},
	), nil
}

func GeneratePrivateKey(encryptType string) (*ecdsa.PrivateKey, error) {
	var priKey *ecdsa.PrivateKey
	var err error
	if encryptType == constants.Prime256v1 {
		priKey, err = NewSecp256r1Key()
	} else if encryptType == constants.Secp256k1 {
		priKey, err = NewSecp256k1Key()
	}

	if err != nil {
		return nil, err
	}
	return priKey, nil
}

func marshalPublicKey(pub *ecdsa.PublicKey) (publicKeyBytes []byte, publicKeyAlgorithm pkix.AlgorithmIdentifier, err error) {
	// Get OID for the curve
	oid, ok := oidFromNamedCurve(pub.Curve)
	if !ok {
		return nil, pkix.AlgorithmIdentifier{}, errors.New("unsupported elliptic curve")
	}

	// Serialize public key
	if pub.Curve == elliptic.P256() {
		// Use ECDH for P-256
		publicDchKey, err := pub.ECDH()
		if err != nil {
			return nil, pkix.AlgorithmIdentifier{}, fmt.Errorf("failed to convert to ECDH key: %v", err)
		}
		publicKeyBytes = publicDchKey.Bytes()
	} else {
		// Manual serialization for other curves (e.g., secp256k1)
		publicKeyBytes = make([]byte, 65)
		publicKeyBytes[0] = 0x04
		xBytes := pub.X.Bytes()
		yBytes := pub.Y.Bytes()
		copy(publicKeyBytes[1:33], xBytes[len(xBytes)-32:])
		copy(publicKeyBytes[33:], yBytes[len(yBytes)-32:])
	}

	// Set algorithm identifier
	publicKeyAlgorithm.Algorithm = oidPublicKeyECDSA
	var paramBytes []byte
	paramBytes, err = asn1.Marshal(oid)
	if err != nil {
		return nil, pkix.AlgorithmIdentifier{}, fmt.Errorf("failed to marshal curve OID: %v", err)
	}
	publicKeyAlgorithm.Parameters.FullBytes = paramBytes

	return publicKeyBytes, publicKeyAlgorithm, nil
}

func NewSecp256r1Key() (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256()
	return ecdsa.GenerateKey(curve, rand.Reader)
}

func NewSecp256k1Key() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(crypto.S256(), rand.Reader)
}

func NewCertificateRequest(name string) *csr.CertificateRequest {
	cr := &csr.CertificateRequest{}
	cr.CN = name
	cr.Names = append(cr.Names, csr.Name{
		OU: "client",
	})

	return cr
}

func PivSKI(key *ecdsa.PrivateKey) []byte {
	if key == nil {
		return nil
	}
	// Marshall the public key
	raw, _ := MarshalPKIXPublicKey(&key.PublicKey)

	// Hash it
	hash := sha256.New()
	hash.Write(raw)
	return hash.Sum(nil)
}

func PubSKI(key *ecdsa.PublicKey) []byte {
	if key == nil {
		return nil
	}

	// Marshall the public key
	raw, _ := MarshalPKIXPublicKey(key)

	// Hash it
	hash := sha256.New()
	hash.Write(raw)

	return hash.Sum(nil)
}

func SHA256Hash(msg []byte) []byte {
	h := sha256.New()
	h.Write(msg)
	hash := h.Sum(nil)

	return hash
}
