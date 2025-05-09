package encryptutil

import "math/big"

type ECDSASignature struct {
	R, S *big.Int
}
