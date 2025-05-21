package node

import (
	"strconv"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type BlockReqData struct {
	base.BaseReqModel
	Body BlockReqDataBody `json:"body"`
}

type BlockReqDataBody struct {
	BlockNumber int64  `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
}

func (f *BlockReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + strconv.FormatInt(f.Body.BlockNumber, 10) + f.Body.BlockHash
}
