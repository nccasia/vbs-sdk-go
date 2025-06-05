package proposal

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-protos-go-apiv2/common"
	"github.com/hyperledger/fabric-protos-go-apiv2/peer"
	"github.com/nccasia/vbs-sdk-go/pkg/common/encrypt"
	"github.com/nccasia/vbs-sdk-go/pkg/core/model/userdata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Proposal struct {
	TxId     string
	Proposal *peer.Proposal
}

type SignedProposal struct {
	TxId                string
	SignedProposalBytes []byte
}

func CreateSignedProposal(channelID, chaincodeName, functionName string, args []string, user *userdata.UserData) (*SignedProposal, error) {
	creator, _ := user.Serialize()
	// Tạo proposal
	proposal, err := createProposal(channelID, chaincodeName, functionName, args, creator)
	if err != nil {
		return nil, fmt.Errorf("failed to create proposal: %w", err)
	}

	// Ký proposal
	signedProposal, err := signProposal(proposal.Proposal, user.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign proposal: %w", err)
	}

	signedProposalBytes, err := proto.Marshal(signedProposal)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signed proposal: %w", err)
	}

	fmt.Printf("signedProposal size: %d bytes\n", len(signedProposalBytes))
	return &SignedProposal{TxId: proposal.TxId, SignedProposalBytes: signedProposalBytes}, nil
}

func SignPayload(data []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := sign(privKey, data)
	if err != nil {
		return nil, err
	}

	envelope := &common.Envelope{Payload: data, Signature: sig}
	envelopeBytes, err := proto.Marshal(envelope)
	if err != nil {
		return nil, err
	}

	return envelopeBytes, nil
}

func createProposal(channelID, chaincodeName, functionName string, args []string, creator []byte) (*Proposal, error) {
	// Tạo nonce random
	nonce := make([]byte, 24)
	_, _ = rand.Read(nonce)

	// 1. Tạo ChaincodeInvocationSpec payload
	ccInput := &peer.ChaincodeInput{
		Args: [][]byte{[]byte(functionName)},
	}
	for _, arg := range args {
		ccInput.Args = append(ccInput.Args, []byte(arg))
	}
	ccSpec := &peer.ChaincodeSpec{
		Type:        peer.ChaincodeSpec_GOLANG,
		ChaincodeId: &peer.ChaincodeID{Name: chaincodeName},
		Input:       ccInput,
	}
	ccInvocationSpec := &peer.ChaincodeInvocationSpec{
		ChaincodeSpec: ccSpec,
	}
	ccInvocationBytes, err := proto.Marshal(ccInvocationSpec)
	if err != nil {
		return nil, err
	}

	ccHdrExt := &peer.ChaincodeHeaderExtension{ChaincodeId: ccSpec.ChaincodeId}
	ccHdrExtBytes, err := proto.Marshal(ccHdrExt)
	if err != nil {
		return nil, err
	}

	// 2. Tạo ChannelHeader (txid, timestamp, type)
	txID := computeTxID(nonce, creator) // Hàm tạo txID
	fmt.Printf("TxId: %s\n", txID)
	chHeader := &common.ChannelHeader{
		Type:      int32(common.HeaderType_ENDORSER_TRANSACTION), // 3
		ChannelId: channelID,
		TxId:      txID,
		Timestamp: timestampNow(),
		Epoch:     0,
		Extension: ccHdrExtBytes,
	}
	chHeaderBytes, err := proto.Marshal(chHeader)
	if err != nil {
		return nil, err
	}

	// 3. Tạo SignatureHeader (creator + nonce)
	sigHeader := &common.SignatureHeader{
		Creator: creator, // certificate identity bytes (serialized Identity proto)
		Nonce:   nonce,
	}
	sigHeaderBytes, err := proto.Marshal(sigHeader)
	if err != nil {
		return nil, err
	}

	// 4. Tạo Header
	header := &common.Header{
		ChannelHeader:   chHeaderBytes,
		SignatureHeader: sigHeaderBytes,
	}
	headerBytes, err := proto.Marshal(header)
	if err != nil {
		return nil, err
	}

	// 5. Tạo Proposal payload (ChaincodeProposalPayload)
	ccProposalPayload := &peer.ChaincodeProposalPayload{
		Input: ccInvocationBytes,
	}
	ccProposalPayloadBytes, err := proto.Marshal(ccProposalPayload)
	if err != nil {
		return nil, err
	}

	// 6. Tạo Proposal
	proposal := &peer.Proposal{
		Header:  headerBytes,
		Payload: ccProposalPayloadBytes,
	}

	return &Proposal{TxId: txID, Proposal: proposal}, nil
}

func timestampNow() *timestamppb.Timestamp {
	now := time.Now()
	return &timestamppb.Timestamp{
		Seconds: int64(now.Unix()),
		Nanos:   int32(now.Nanosecond()),
	}
}

func computeTxID(nonce, creator []byte) string {
	data := append(nonce, creator...)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func signProposal(proposal *peer.Proposal, privKey *ecdsa.PrivateKey) (*peer.SignedProposal, error) {
	proposalBytes, err := proto.Marshal(proposal)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal proposal: %w", err)
	}

	sig, err := sign(privKey, proposalBytes)
	if err != nil {
		return nil, err
	}

	return &peer.SignedProposal{ProposalBytes: proposalBytes, Signature: sig}, nil
}

func sign(privKey *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	digest := sha256.Sum256(data)
	sig, err := encrypt.SignECDSA(privKey, digest[:])
	if err != nil {
		return nil, err
	}

	return sig, nil
}
