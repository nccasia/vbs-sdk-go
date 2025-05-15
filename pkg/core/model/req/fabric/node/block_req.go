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
	BlockNumber uint64 `json:"blockNumber"`
	BlockHash   string `json:"blockHash,optional"`
	TxId        string `json:"txId,optional"`
	DataType    string `json:"dataType,optional"`
}

func (f *BlockReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()
	fp = fp + strconv.FormatUint(f.Body.BlockNumber, 10)
	fp = fp + f.Body.BlockHash
	fp = fp + f.Body.TxId
	fp = fp + f.Body.DataType
	return fp
}
