package node

import (
	"strconv"

	"github.com/nccasia/vbs-sdk-go/pkg/core/model/base"
)

type SystemResData struct {
	base.BaseResModel
	Body *SystemResDataBody `json:"body"`
}

type SystemResDataBody struct {
	ChainID        string     `json:"chainId"`
	BlockNumber    uint64     `json:"blockNumber"`
	TxCount        uint64     `json:"txCount"`
	PendingTxCount uint64     `json:"pendingTxCount"`
	NodeCount      int        `json:"nodeCount"`
	Nodes          []NodeInfo `json:"nodes"`
}

type NodeInfo struct {
	NodeID   string `json:"nodeId"`
	NodeType string `json:"nodeType"`
}

func (f *SystemResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.ChainID
	fp = fp + strconv.FormatUint(f.Body.BlockNumber, 10)
	fp = fp + strconv.FormatUint(f.Body.TxCount, 10)
	fp = fp + strconv.FormatUint(f.Body.PendingTxCount, 10)
	fp = fp + strconv.Itoa(f.Body.NodeCount)

	for _, tx := range f.Body.Nodes {
		fp = fp + tx.NodeID
		fp = fp + tx.NodeType
	}

	return fp
}
