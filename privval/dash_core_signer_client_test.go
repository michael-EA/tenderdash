package privval

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/tendermint/tendermint/types"
	"testing"
)

// TestDashCoreVoteSigning will create a vote and attempt to sign it.
func TestDashCoreVoteSigning(t *testing.T) {
	validatorProTxHash, err := hex.DecodeString("f4406d978ae15b97976411212e3abd66515a285d901ace06758dc1012030da08")
	require.NoError(t, err)
	blockHash, err := hex.DecodeString("a3406d978ae15b97976411212e3abd44515a285d932ace06758dc1012030da45")
	require.NoError(t, err)
	partSetHeaderHash, err := hex.DecodeString("56706d978ae15b97976411212e3abd45555a285d932ace06758dc1012030d980")
	require.NoError(t, err)
	lastAppHash, err := hex.DecodeString("513b4e704ee33b376cb2049f966d93ad43058039a3e0e089bc58e3a8caf80adc")
	require.NoError(t, err)
	quorumHash, err := hex.DecodeString("c1f878a748532972813515cdd238ce7e57603c998b1ebe2af50947c19e010000")
	require.NoError(t, err)
	expectedVoteStateSignBytesString := "2904000000000000513b4e704ee33b376cb2049f966d93ad43058039a3e0e089bc58e3a8caf80adc"
	expectedVoteStateRequestIDString := "5860d471bd8dde139cbc4d382003c3ecfc3f6d17853602d6541392bdeaa19aa6"
	expectedVoteStateSignIDString := "8e6af40613e1d8c67076e54880e6ac5631dc749f319b7ff0390b09dd625e733f"
	vote := &types.Vote{
		ValidatorProTxHash: validatorProTxHash,
		ValidatorIndex:     0,
		Height:             1090,
		Round:              0,
		Type:               tmproto.PrecommitType,
		BlockID:            types.BlockID{Hash: blockHash, PartSetHeader: types.PartSetHeader{
			Total: 4,
			Hash: partSetHeaderHash,
		}},
		StateID:            types.StateID{LastAppHash: lastAppHash},
	}
	protoVote := vote.ToProto()
	voteStateSignBytes := types.VoteStateSignBytes("test", protoVote)
	require.EqualValues(t, hex.EncodeToString(voteStateSignBytes), expectedVoteStateSignBytesString)

	stateRequestID := types.VoteStateRequestIDProto(protoVote)
	require.EqualValues(t, hex.EncodeToString(stateRequestID), expectedVoteStateRequestIDString)

	stateSignID :=types.VoteStateSignID("test", protoVote, 101, quorumHash)
	require.EqualValues(t, hex.EncodeToString(stateSignID), expectedVoteStateSignIDString)
}
