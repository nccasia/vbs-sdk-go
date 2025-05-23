package node

import (
	"strconv"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type BlockResData struct {
	base.BaseResModel
	Body *BlockResDataBody `json:"body"`
}

type BlockResDataBody struct {
	BlockHash       string                        `json:"blockHash"`
	BlockNumber     int64                         `json:"blockNumber"`
	ParentBlockHash string                        `json:"parentBlockHash"`
	BlockSize       int64                         `json:"blockSize"`
	BlockTime       int64                         `json:"blockTime"`
	Author          string                        `json:"author"`
	Transactions    []*TransactionInfoResDataBody `json:"transactions"`
}

func (f *BlockResDataBody) GetEncryptionValue() string {
	fb := ""
	fb = fb + f.BlockHash
	fb = fb + strconv.FormatInt(f.BlockNumber, 10)
	fb = fb + f.ParentBlockHash
	fb = fb + strconv.FormatInt(f.BlockSize, 10)
	fb = fb + strconv.FormatInt(f.BlockTime, 10)
	fb = fb + f.Author

	for _, tx := range f.Transactions {
		fb = fb + tx.GetEncryptionValue()
	}

	return fb
}

func (f *BlockResData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.GetEncryptionValue()
}
