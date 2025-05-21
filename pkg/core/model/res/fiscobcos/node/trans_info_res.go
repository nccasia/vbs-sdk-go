package node

import (
	"strconv"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type TransactionInfoResData struct {
	base.BaseResModel
	Body *TransactionInfoResDataBody `json:"body"`
}

type TransactionInfoResDataBody struct {
	TxId        string `json:"txId"`
	BlockHash   string `json:"blockHash"`
	BlockNumber uint64 `json:"blockNumber"`
	GasUsed     int64  `json:"gasUsed"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       int64  `json:"value"`
	Input       string `json:"input"`
}

func (f *TransactionInfoResDataBody) GetEncryptionValue() string {
	fb := ""
	fb = fb + f.TxId
	fb = fb + f.BlockHash
	fb = fb + strconv.FormatUint(f.BlockNumber, 10)
	fb = fb + strconv.FormatInt(f.GasUsed, 10)
	fb = fb + f.From
	fb = fb + f.To
	fb = fb + strconv.FormatInt(f.Value, 10)
	fb = fb + f.Input
	return fb
}

func (f *TransactionInfoResData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.GetEncryptionValue()
}
