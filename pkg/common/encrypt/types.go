package encrypt

import "math/big"

type ECDSASignature struct {
	R, S *big.Int
}
