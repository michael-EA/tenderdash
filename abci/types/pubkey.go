package types

import (
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/bls12381"
	cryptoenc "github.com/tendermint/tendermint/crypto/encoding"
	crypto2 "github.com/tendermint/tendermint/proto/tendermint/crypto"
	pbtypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

func UpdateValidator(
	proTxHash crypto.ProTxHash, pubkeyBytes []byte, power int64, ip pbtypes.IPAddress, port uint16) ValidatorUpdate {
	valUpdate := ValidatorUpdate{
		Power:     power,
		ProTxHash: proTxHash,
		Address:   &pbtypes.NetworkEndpoint{IP: &ip, Port: uint32(port)},
	}

	if len(pubkeyBytes) > 0 {
		pke := bls12381.PubKey(pubkeyBytes)
		pkp, err := cryptoenc.PubKeyToProto(pke)
		if err != nil {
			panic(err)
		}
		valUpdate.PubKey = &pkp
	}
	return valUpdate
}

func UpdateValidatorSet(validatorUpdates []ValidatorUpdate, thresholdPublicKey crypto2.PublicKey) ValidatorSetUpdate {
	return ValidatorSetUpdate{
		ValidatorUpdates:   validatorUpdates,
		ThresholdPublicKey: thresholdPublicKey,
	}
}

func UpdateThresholdPublicKey(pk []byte) ThresholdPublicKeyUpdate {
	pke := bls12381.PubKey(pk)
	pkp, err := cryptoenc.PubKeyToProto(pke)
	if err != nil {
		panic(err)
	}

	return ThresholdPublicKeyUpdate{
		ThresholdPublicKey: pkp,
	}
}

func UpdateQuorumHash(quorumHash crypto.QuorumHash) QuorumHashUpdate {
	return QuorumHashUpdate{
		QuorumHash: quorumHash,
	}
}
