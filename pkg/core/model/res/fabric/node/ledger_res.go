package node

import (
	"strconv"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type LedgerResData struct {
	base.BaseResModel
	Body *LedgerResDataBody `json:"body"`
}

type LedgerResDataBody struct {
	BlockHash    string `json:"blockHash"`
	PreBlockHash string `json:"preBlockHash"`
	Height       uint64 `json:"height"`
}

func (f *LedgerResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.BlockHash
	fp = fp + strconv.FormatUint(f.Body.Height, 10)
	fp = fp + f.Body.PreBlockHash

	return fp
}
