package types

import (
	"github.com/dashevo/dashd-go/btcjson"
	"github.com/tendermint/tendermint/crypto/bls12381"
)

var (
	SignatureSize = bls12381.SignatureSize
)

// Signable is an interface for all signable things.
// It typically removes signatures before serializing.
// SignBytes returns the bytes to be signed
// NOTE: chainIDs are part of the SignBytes but not
// necessarily the object themselves.
// NOTE: Expected to panic if there is an error marshalling.
type Signable interface {
	SignBytes(chainID string) []byte
	SignRequestID() []byte
	SignID(chainID string, quorumType btcjson.LLMQType, quorumHash []byte) []byte
}
