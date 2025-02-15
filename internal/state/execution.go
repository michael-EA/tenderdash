package state

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	abciclient "github.com/tendermint/tendermint/abci/client"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	cryptoenc "github.com/tendermint/tendermint/crypto/encoding"
	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/internal/eventbus"
	"github.com/tendermint/tendermint/internal/mempool"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/libs/log"
	tmstate "github.com/tendermint/tendermint/proto/tendermint/state"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/tendermint/tendermint/types"
)

//-----------------------------------------------------------------------------
// BlockExecutor handles block execution and state updates.
// It exposes ApplyBlock(), which validates & executes the block, updates state w/ ABCI responses,
// then commits and updates the mempool atomically, then saves state.

// BlockExecutor provides the context and accessories for properly executing a block.
type BlockExecutor struct {
	// save state, validators, consensus params, abci responses here
	store Store

	// use blockstore for the pruning functions.
	blockStore BlockStore

	// execute the app against this
	appClient abciclient.Client

	// events
	eventBus types.BlockEventPublisher

	// manage the mempool lock during commit
	// and update both with block results after commit.
	mempool mempool.Mempool
	evpool  EvidencePool

	logger  log.Logger
	metrics *Metrics

	appHashSize int

	// cache the verification results over a single height
	cache map[string]struct{}

	// the next core chain lock that we can propose
	NextCoreChainLock *types.CoreChainLock
}

// NewBlockExecutor returns a new BlockExecutor with the passed-in EventBus.
func NewBlockExecutor(
	stateStore Store,
	logger log.Logger,
	appClient abciclient.Client,
	pool mempool.Mempool,
	evpool EvidencePool,
	blockStore BlockStore,
	eventBus *eventbus.EventBus,
	metrics *Metrics,
) *BlockExecutor {
	return &BlockExecutor{
		eventBus:    eventBus,
		store:       stateStore,
		appClient:   appClient,
		mempool:     pool,
		evpool:      evpool,
		logger:      logger,
		metrics:     metrics,
		cache:       make(map[string]struct{}),
		blockStore:  blockStore,
		appHashSize: 32, // TODO change on constant
	}
}

func (blockExec *BlockExecutor) Store() Store {
	return blockExec.store
}

// CreateProposalBlock calls state.MakeBlock with evidence from the evpool
// and txs from the mempool. The max bytes must be big enough to fit the commit.
// Up to 1/10th of the block space is allocated for maximum sized evidence.
// The rest is given to txs, up to the max gas.
//
// Contract: application will not return more bytes than are sent over the wire.
func (blockExec *BlockExecutor) CreateProposalBlock(
	ctx context.Context,
	height int64,
	state State,
	commit *types.Commit,
	proposerProTxHash []byte,
	proposedAppVersion uint64,
) (*types.Block, error) {

	maxBytes := state.ConsensusParams.Block.MaxBytes
	maxGas := state.ConsensusParams.Block.MaxGas

	evidence, evSize := blockExec.evpool.PendingEvidence(state.ConsensusParams.Evidence.MaxBytes)

	// Fetch a limited amount of valid txs
	maxDataBytes := types.MaxDataBytes(maxBytes, crypto.BLS12381, evSize, state.Validators.Size())

	nextCoreChainLock := blockExec.NextCoreChainLock
	if nextCoreChainLock != nil &&
		nextCoreChainLock.CoreBlockHeight <= state.LastCoreChainLockedBlockHeight {
		nextCoreChainLock = nil
	}

	// Pass proposed app version only if it's higher than current network app version
	if proposedAppVersion <= state.Version.Consensus.App {
		proposedAppVersion = 0
	}

	txs := blockExec.mempool.ReapMaxBytesMaxGas(maxDataBytes, maxGas)
	block := state.MakeBlock(height, nextCoreChainLock, txs, commit, evidence, proposerProTxHash, proposedAppVersion)

	localLastCommit := buildLastCommitInfo(block, blockExec.store, state.InitialHeight)
	version := block.Version.ToProto()
	rpp, err := blockExec.appClient.PrepareProposal(
		ctx,
		&abci.RequestPrepareProposal{
			MaxTxBytes:          maxDataBytes,
			Txs:                 block.Txs.ToSliceOfBytes(),
			LocalLastCommit:     abci.ExtendedCommitInfo(localLastCommit),
			ByzantineValidators: block.Evidence.ToABCI(),
			Height:              block.Height,
			Time:                block.Time,
			NextValidatorsHash:  block.NextValidatorsHash,

			// Dash's fields
			CoreChainLockedHeight: block.CoreChainLockedHeight,
			ProposerProTxHash:     block.ProposerProTxHash,
			ProposedAppVersion:    block.ProposedAppVersion,
			Version:               &version,
		},
	)
	if err != nil {
		// The App MUST ensure that only valid (and hence 'processable') transactions
		// enter the mempool. Hence, at this point, we can't have any non-processable
		// transaction causing an error.
		//
		// Also, the App can simply skip any transaction that could cause any kind of trouble.
		// Either way, we cannot recover in a meaningful way, unless we skip proposing
		// this block, repair what caused the error and try again. Hence, we return an
		// error for now (the production code calling this function is expected to panic).
		return nil, err
	}
	txrSet := types.NewTxRecordSet(rpp.TxRecords)

	if err := txrSet.Validate(maxDataBytes, block.Txs); err != nil {
		return nil, err
	}

	for _, rtx := range txrSet.RemovedTxs() {
		if err := blockExec.mempool.RemoveTxByKey(rtx.Key()); err != nil {
			blockExec.logger.Debug("error removing transaction from the mempool", "error", err, "tx hash", rtx.Hash())
		}
	}
	itxs := txrSet.IncludedTxs()

	return state.MakeBlock(
		height,
		nextCoreChainLock,
		itxs,
		commit,
		evidence,
		proposerProTxHash,
		proposedAppVersion,
	), nil
}

func (blockExec *BlockExecutor) ProcessProposal(
	ctx context.Context,
	block *types.Block,
	state State,
) (bool, error) {
	resp, err := blockExec.appClient.ProcessProposal(ctx, &abci.RequestProcessProposal{
		Hash:                block.Header.Hash(),
		Height:              block.Header.Height,
		Time:                block.Header.Time,
		Txs:                 block.Data.Txs.ToSliceOfBytes(),
		ProposedLastCommit:  buildLastCommitInfo(block, blockExec.store, state.InitialHeight),
		ByzantineValidators: block.Evidence.ToABCI(),
		ProposerProTxHash:   block.ProposerProTxHash,
		NextValidatorsHash:  block.NextValidatorsHash,
	})
	if err != nil {
		return false, ErrInvalidBlock(err)
	}
	if resp.IsStatusUnknown() {
		panic(fmt.Sprintf("ProcessProposal responded with status %s", resp.Status.String()))
	}

	return resp.IsAccepted(), nil
}

// ValidateBlock validates the given block against the given state.
// If the block is invalid, it returns an error.
// Validation does not mutate state, but does require historical information from the stateDB,
// ie. to verify evidence from a validator at an old height.
func (blockExec *BlockExecutor) ValidateBlock(ctx context.Context, state State, block *types.Block) error {
	hash := block.Hash()
	if _, ok := blockExec.cache[hash.String()]; ok {
		return nil
	}

	err := validateBlock(state, block)
	if err != nil {
		return err
	}

	err = blockExec.evpool.CheckEvidence(ctx, block.Evidence)
	if err != nil {
		return err
	}

	blockExec.cache[hash.String()] = struct{}{}
	return nil
}

// ValidateBlockChainLock validates the given block chain lock against the given state.
// If the block is invalid, it returns an error.
// Validation does not mutate state, but does require historical information from the stateDB,
// ie. to verify evidence from a validator at an old height.
func (blockExec *BlockExecutor) ValidateBlockChainLock(ctx context.Context, state State, block *types.Block) error {
	return validateBlockChainLock(ctx, blockExec.appClient, state, block)
}

// ValidateBlockTime validates the given block time against the given state.
// If the block is invalid, it returns an error.
// Validation does not mutate state, but does require historical information from the stateDB,
// ie. to verify evidence from a validator at an old height.
func (blockExec *BlockExecutor) ValidateBlockTime(
	allowedTimeWindow time.Duration,
	state State,
	block *types.Block,
) error {
	return validateBlockTime(allowedTimeWindow, state, block)
}

// ApplyBlock validates the block against the state, executes it against the app,
// fires the relevant events, commits the app, and saves the new state and responses.
// It returns the new state.
// It's the only function that needs to be called
// from outside this package to process and commit an entire block.
// It takes a blockID to avoid recomputing the parts hash.
func (blockExec *BlockExecutor) ApplyBlock(
	ctx context.Context,
	state State,
	proTxHash crypto.ProTxHash,
	blockID types.BlockID, block *types.Block,
) (State, error) {

	// validate the block if we haven't already
	if err := blockExec.ValidateBlock(ctx, state, block); err != nil {
		return state, ErrInvalidBlock(err)
	}
	startTime := time.Now().UnixNano()
	txs := block.Txs.ToSliceOfBytes()
	version := block.Header.Version.ToProto()
	finalizeBlockResponse, err := blockExec.appClient.FinalizeBlock(
		ctx,
		&abci.RequestFinalizeBlock{
			Hash:                block.Hash(),
			Height:              block.Header.Height,
			Time:                block.Header.Time,
			Txs:                 txs,
			DecidedLastCommit:   buildLastCommitInfo(block, blockExec.store, state.InitialHeight),
			ByzantineValidators: block.Evidence.ToABCI(),
			NextValidatorsHash:  block.NextValidatorsHash,

			// Dash's fields
			ProposerProTxHash:     block.ProposerProTxHash,
			CoreChainLockedHeight: block.CoreChainLockedHeight,
			ProposedAppVersion:    block.ProposedAppVersion,
			Version:               &version,
		},
	)
	endTime := time.Now().UnixNano()
	blockExec.metrics.BlockProcessingTime.Observe(float64(endTime-startTime) / 1000000)
	if err != nil {
		return state, ErrProxyAppConn(err)
	}

	abciResponses := &tmstate.ABCIResponses{
		FinalizeBlock: finalizeBlockResponse,
	}

	// Save the results before we commit.
	if err := blockExec.store.SaveABCIResponses(block.Height, abciResponses); err != nil {
		return state, err
	}

	// validate the validator updates and convert to tendermint types
	err = validateValidatorSetUpdate(finalizeBlockResponse.ValidatorSetUpdate, state.ConsensusParams.Validator)
	if err != nil {
		return state, fmt.Errorf("error in validator updates: %w", err)
	}

	nextCoreChainLock, err := types.CoreChainLockFromProto(
		finalizeBlockResponse.NextCoreChainLockUpdate,
	)
	if err != nil {
		return state, fmt.Errorf("error in chain lock from proto: %v", err)
	}

	// The quorum type should not even matter here
	validators, thresholdPublicKey, quorumHash, err := types.PB2TM.ValidatorUpdatesFromValidatorSet(finalizeBlockResponse.ValidatorSetUpdate)
	if err != nil {
		return state, fmt.Errorf("error in chain lock from proto: %v", err)
	}

	if len(validators) > 0 {
		blockExec.logger.Debug(
			"updates to validators",
			"height", block.Height,
			"quorumHash", quorumHash.ShortString(),
			"thresholdPublicKey", tmbytes.HexBytes(thresholdPublicKey.Bytes()).ShortString(),
			"updates", types.ValidatorListString(validators),
		)
	}

	// Update the state with the block and responses.
	rs, err := abci.MarshalTxResults(finalizeBlockResponse.TxResults)
	if err != nil {
		return state, fmt.Errorf("marshaling TxResults: %w", err)
	}
	h := merkle.HashFromByteSlices(rs)
	state, err = state.Update(proTxHash, blockID, &block.Header, h, finalizeBlockResponse.ConsensusParamUpdates, validators, thresholdPublicKey, quorumHash)
	if err != nil {
		return state, fmt.Errorf("commit failed for application: %w", err)
	}

	// Lock mempool, commit app state, update mempoool.
	appHash, retainHeight, err := blockExec.Commit(ctx, state, block, finalizeBlockResponse.TxResults)
	if err != nil {
		return state, fmt.Errorf("commit failed for application: %w", err)
	}

	// Update evpool with the latest state.
	blockExec.evpool.Update(ctx, state, block.Evidence)

	// Update the next core chain lock that we can propose
	blockExec.NextCoreChainLock = nextCoreChainLock

	// Update the app hash and save the state.
	state.AppHash = appHash
	if err := blockExec.store.Save(state); err != nil {
		return state, err
	}

	// Prune old heights, if requested by ABCI app.
	if retainHeight > 0 {
		pruned, err := blockExec.pruneBlocks(retainHeight)
		if err != nil {
			blockExec.logger.Error("failed to prune blocks", "retain_height", retainHeight, "err", err)
		} else {
			blockExec.logger.Debug("pruned blocks", "pruned", pruned, "retain_height", retainHeight)
		}
	}

	// reset the verification cache
	blockExec.cache = make(map[string]struct{})

	// Events are fired after everything else.
	// NOTE: if we crash between Commit and Save, events wont be fired during replay
	fireEvents(blockExec.logger, blockExec.eventBus, block, blockID, finalizeBlockResponse, state.NextValidators)

	return state, nil
}

func (blockExec *BlockExecutor) ExtendVote(ctx context.Context, vote *types.Vote) ([]*abci.ExtendVoteExtension, error) {
	resp, err := blockExec.appClient.ExtendVote(ctx, &abci.RequestExtendVote{
		Hash:   vote.BlockID.Hash,
		Height: vote.Height,
	})
	if err != nil {
		panic(fmt.Errorf("ExtendVote call failed: %w", err))
	}
	return resp.VoteExtensions, nil
}

func (blockExec *BlockExecutor) VerifyVoteExtension(ctx context.Context, vote *types.Vote) error {
	var extensions []*abci.ExtendVoteExtension
	if vote.VoteExtensions != nil {
		extensions = vote.VoteExtensions.ToExtendProto()
	}

	resp, err := blockExec.appClient.VerifyVoteExtension(ctx, &abci.RequestVerifyVoteExtension{
		Hash:               vote.BlockID.Hash,
		Height:             vote.Height,
		ValidatorProTxHash: vote.ValidatorProTxHash,
		VoteExtensions:     extensions,
	})
	if err != nil {
		panic(fmt.Errorf("VerifyVoteExtension call failed: %w", err))
	}

	if !resp.IsOK() {
		return types.ErrVoteInvalidExtension
	}

	return nil
}

// Commit locks the mempool, runs the ABCI Commit message, and updates the
// mempool.
// It returns the result of calling abci.Commit (the AppHash) and the height to retain (if any).
// The Mempool must be locked during commit and update because state is
// typically reset on Commit and old txs must be replayed against committed
// state before new txs are run in the mempool, lest they be invalid.
func (blockExec *BlockExecutor) Commit(
	ctx context.Context,
	state State,
	block *types.Block,
	txResults []*abci.ExecTxResult,
) ([]byte, int64, error) {
	blockExec.mempool.Lock()
	defer blockExec.mempool.Unlock()

	// while mempool is Locked, flush to ensure all async requests have completed
	// in the ABCI app before Commit.
	err := blockExec.mempool.FlushAppConn(ctx)
	if err != nil {
		blockExec.logger.Error("client error during mempool.FlushAppConn", "err", err)
		return nil, 0, err
	}

	// Commit block, get hash back
	res, err := blockExec.appClient.Commit(ctx)
	if err != nil {
		blockExec.logger.Error("client error during proxyAppConn.Commit", "err", err)
		return nil, 0, err
	}

	// we force the abci app to return only 32 byte app hashes (set to 20 temporarily)
	if res.Data != nil && len(res.Data) != blockExec.appHashSize {
		blockExec.logger.Error(
			"Client returned invalid app hash size", "bytesLength", len(res.Data),
		)
		return nil, 0, errors.New("invalid App Hash size")
	}

	// ResponseCommit has no error code - just data
	blockExec.logger.Info(
		"committed state",
		"height", block.Height,
		"core_height", block.CoreChainLockedHeight,
		"num_txs", len(block.Txs),
		"app_hash", fmt.Sprintf("%X", res.Data),
	)

	// Update mempool.
	err = blockExec.mempool.Update(
		ctx,
		block.Height,
		block.Txs,
		txResults,
		TxPreCheckForState(state),
		TxPostCheckForState(state),
	)

	return res.Data, res.RetainHeight, err
}

func buildLastCommitInfo(block *types.Block, store Store, initialHeight int64) abci.CommitInfo {
	if block.Height == initialHeight {
		// there is no last commit for the initial height.
		// return an empty value.
		return abci.CommitInfo{}
	}
	return abci.CommitInfo{
		Round:                   block.LastCommit.Round,
		QuorumHash:              block.LastCommit.QuorumHash,
		BlockSignature:          block.LastCommit.ThresholdBlockSignature,
		StateSignature:          block.LastCommit.ThresholdStateSignature,
		ThresholdVoteExtensions: types.ThresholdExtensionSignToProto(block.LastCommit.ThresholdVoteExtensions),
	}
}

func validateValidatorSetUpdate(
	abciValidatorSetUpdate *abci.ValidatorSetUpdate,
	params types.ValidatorParams,
) error {
	// if there was no update return no error
	if abciValidatorSetUpdate == nil {
		return nil
	}
	if len(abciValidatorSetUpdate.ValidatorUpdates) != 0 &&
		abciValidatorSetUpdate.ThresholdPublicKey.Sum == nil {
		return fmt.Errorf("received validator updates without a threshold public key")
	}
	return validateValidatorUpdates(abciValidatorSetUpdate.ValidatorUpdates, params)
}

func validateValidatorUpdates(abciUpdates []abci.ValidatorUpdate, params types.ValidatorParams) error {
	for _, valUpdate := range abciUpdates {
		if valUpdate.GetPower() < 0 {
			return fmt.Errorf("voting power can't be negative %v", valUpdate)
		} else if valUpdate.GetPower() == 0 {
			// continue, since this is deleting the validator, and thus there is no
			// pubkey to check
			continue
		}

		// Check if validator's pubkey matches an ABCI type in the consensus params
		if valUpdate.PubKey != nil {
			pk, err := cryptoenc.PubKeyFromProto(*valUpdate.PubKey)
			if err != nil {
				return err
			}
			if !params.IsValidPubkeyType(pk.Type()) {
				return fmt.Errorf(
					"validator %v is using pubkey %s, which is unsupported for consensus",
					valUpdate,
					pk.Type(),
				)
			}
			if err := validatePubKey(pk); err != nil {
				return fmt.Errorf("public key in validator %X is invalid: %w", valUpdate.ProTxHash, err)
			}
		}

		if valUpdate.ProTxHash == nil {
			return fmt.Errorf(
				"validator %v does not have a protxhash, which is needed for consensus",
				valUpdate,
			)
		}

		if len(valUpdate.ProTxHash) != crypto.ProTxHashSize {
			return fmt.Errorf(
				"validator %v is using protxhash %s, which is not the required length",
				valUpdate,
				valUpdate.ProTxHash,
			)
		}

		if valUpdate.NodeAddress != "" {
			_, err := types.ParseValidatorAddress(valUpdate.NodeAddress)
			if err != nil {
				return fmt.Errorf("cannot parse validator address %s: %w", valUpdate.NodeAddress, err)
			}
		}
	}
	return nil
}

// Update returns a copy of state with the fields set using the arguments passed in.
func (state State) Update(
	nodeProTxHash crypto.ProTxHash,
	blockID types.BlockID,
	header *types.Header,
	resultsHash []byte,
	consensusParamUpdates *tmtypes.ConsensusParams,
	validatorUpdates []*types.Validator,
	newThresholdPublicKey crypto.PubKey,
	quorumHash crypto.QuorumHash,
) (State, error) {

	// Copy the valset so we can apply changes from FinalizeBlock
	// and update s.LastValidators and s.Validators.
	nValSet := state.NextValidators.Copy()

	// We need to generate LastStateID before changing the state
	lastStateID := state.StateID()

	// Update the validator set with the latest abciResponses.
	lastHeightValsChanged := state.LastHeightValidatorsChanged
	if len(validatorUpdates) > 0 {
		if bytes.Equal(nValSet.QuorumHash, quorumHash) {
			err := nValSet.UpdateWithChangeSet(validatorUpdates, newThresholdPublicKey, quorumHash)
			if err != nil {
				return state, fmt.Errorf("error changing validator set: %w", err)
			}
			// Change results from this height but only applies to the next next height.
			lastHeightValsChanged = header.Height + 1 + 1
		} else {
			nValSet = types.NewValidatorSetWithLocalNodeProTxHash(validatorUpdates, newThresholdPublicKey,
				state.Validators.QuorumType, quorumHash, nodeProTxHash)
			// Change results from this height but only applies to the next next height.
			lastHeightValsChanged = header.Height + 1 + 1
		}
	}

	// Update validator proposer priority and set state variables.
	nValSet.IncrementProposerPriority(1)

	// Update the params with the latest abciResponses.
	nextParams := state.ConsensusParams
	lastHeightParamsChanged := state.LastHeightConsensusParamsChanged
	if consensusParamUpdates != nil {
		// NOTE: must not mutate state.ConsensusParams
		nextParams = state.ConsensusParams.UpdateConsensusParams(consensusParamUpdates)
		err := nextParams.ValidateConsensusParams()
		if err != nil {
			return state, fmt.Errorf("error updating consensus params: %w", err)
		}

		state.Version.Consensus.App = nextParams.Version.AppVersion

		// Change results from this height but only applies to the next height.
		lastHeightParamsChanged = header.Height + 1
	}

	nextVersion := state.Version

	// NOTE: the AppHash and the VoteExtension has not been populated.
	// It will be filled on state.Save.
	return State{
		Version:                          nextVersion,
		ChainID:                          state.ChainID,
		InitialHeight:                    state.InitialHeight,
		LastBlockHeight:                  header.Height,
		LastBlockID:                      blockID,
		LastStateID:                      lastStateID,
		LastBlockTime:                    header.Time,
		LastCoreChainLockedBlockHeight:   header.CoreChainLockedHeight,
		NextValidators:                   nValSet,
		Validators:                       state.NextValidators.Copy(),
		LastValidators:                   state.Validators.Copy(),
		LastHeightValidatorsChanged:      lastHeightValsChanged,
		ConsensusParams:                  nextParams,
		LastHeightConsensusParamsChanged: lastHeightParamsChanged,
		LastResultsHash:                  resultsHash,
		AppHash:                          nil,
	}, nil
}

// SetNextCoreChainLock ...
func (blockExec *BlockExecutor) SetNextCoreChainLock(coreChainLock *types.CoreChainLock) {
	blockExec.NextCoreChainLock = coreChainLock
}

// SetAppHashSize ...
func (blockExec *BlockExecutor) SetAppHashSize(size int) {
	blockExec.appHashSize = size
}

// Fire NewBlock, NewBlockHeader.
// Fire TxEvent for every tx.
// NOTE: if Tendermint crashes before commit, some or all of these events may be published again.
func fireEvents(
	logger log.Logger,
	eventBus types.BlockEventPublisher,
	block *types.Block,
	blockID types.BlockID,
	finalizeBlockResponse *abci.ResponseFinalizeBlock,
	validatorSetUpdate *types.ValidatorSet,
) {
	if err := eventBus.PublishEventNewBlock(types.EventDataNewBlock{
		Block:               block,
		BlockID:             blockID,
		ResultFinalizeBlock: *finalizeBlockResponse,
	}); err != nil {
		logger.Error("failed publishing new block", "err", err)
	}

	if err := eventBus.PublishEventNewBlockHeader(types.EventDataNewBlockHeader{
		Header:              block.Header,
		NumTxs:              int64(len(block.Txs)),
		ResultFinalizeBlock: *finalizeBlockResponse,
	}); err != nil {
		logger.Error("failed publishing new block header", "err", err)
	}

	if len(block.Evidence) != 0 {
		for _, ev := range block.Evidence {
			if err := eventBus.PublishEventNewEvidence(types.EventDataNewEvidence{
				Evidence: ev,
				Height:   block.Height,
			}); err != nil {
				logger.Error("failed publishing new evidence", "err", err)
			}
		}
	}

	// sanity check
	if len(finalizeBlockResponse.TxResults) != len(block.Data.Txs) {
		panic(fmt.Sprintf("number of TXs (%d) and ABCI TX responses (%d) do not match",
			len(block.Data.Txs), len(finalizeBlockResponse.TxResults)))
	}

	for i, tx := range block.Data.Txs {
		if err := eventBus.PublishEventTx(types.EventDataTx{
			TxResult: abci.TxResult{
				Height: block.Height,
				Index:  uint32(i),
				Tx:     tx,
				Result: *(finalizeBlockResponse.TxResults[i]),
			},
		}); err != nil {
			logger.Error("failed publishing event TX", "err", err)
		}
	}

	if validatorSetUpdate != nil {
		if err := eventBus.PublishEventValidatorSetUpdates(
			types.EventDataValidatorSetUpdate{
				ValidatorSetUpdates: validatorSetUpdate.Validators,
				ThresholdPublicKey:  validatorSetUpdate.ThresholdPublicKey,
				QuorumHash:          validatorSetUpdate.QuorumHash,
			}); err != nil {
			logger.Error("failed publishing event validator-set update", "err", err)
		}
	}
}

//----------------------------------------------------------------------------------------------------
// Execute block without state. TODO: eliminate

// ExecCommitBlock executes and commits a block on the proxyApp without validating or mutating the state.
// It returns the application root hash (result of abci.Commit).
func ExecCommitBlock(
	ctx context.Context,
	be *BlockExecutor,
	appConn abciclient.Client,
	block *types.Block,
	logger log.Logger,
	store Store,
	initialHeight int64,
	s State,
) ([]byte, error) {
	version := block.Header.Version.ToProto()
	finalizeBlockResponse, err := appConn.FinalizeBlock(
		ctx,
		&abci.RequestFinalizeBlock{
			Hash:                block.Hash(),
			Height:              block.Height,
			Time:                block.Time,
			Txs:                 block.Txs.ToSliceOfBytes(),
			DecidedLastCommit:   buildLastCommitInfo(block, store, initialHeight),
			ByzantineValidators: block.Evidence.ToABCI(),

			// Dash's fields
			CoreChainLockedHeight: block.CoreChainLockedHeight,
			ProposerProTxHash:     block.ProposerProTxHash,
			ProposedAppVersion:    block.ProposedAppVersion,
			Version:               &version,
		},
	)

	if err != nil {
		logger.Error("executing block", "err", err)
		return nil, err
	}
	logger.Info("executed block", "height", block.Height)

	// the BlockExecutor condition is using for the final block replay process.
	if be != nil {
		err = validateValidatorSetUpdate(finalizeBlockResponse.ValidatorSetUpdate, s.ConsensusParams.Validator)
		if err != nil {
			logger.Error("validating validator updates", "err", err)
			return nil, err
		}

		validatorUpdates, thresholdPublicKeyUpdate, quorumHash, err :=
			types.PB2TM.ValidatorUpdatesFromValidatorSet(finalizeBlockResponse.ValidatorSetUpdate)
		if err != nil {
			logger.Error("converting validator updates to native types", "err", err)
			return nil, err
		}

		validatorSetUpdate := s.Validators.Copy()
		err = validatorSetUpdate.UpdateWithChangeSet(validatorUpdates, thresholdPublicKeyUpdate, quorumHash)
		if err != nil {
			return nil, err
		}

		bps, err := block.MakePartSet(types.BlockPartSizeBytes)
		if err != nil {
			return nil, err
		}

		blockID := types.BlockID{Hash: block.Hash(), PartSetHeader: bps.Header()}
		fireEvents(be.logger, be.eventBus, block, blockID, finalizeBlockResponse, validatorSetUpdate)
	}

	// Commit block, get hash back
	res, err := appConn.Commit(ctx)
	if err != nil {
		logger.Error("client error during proxyAppConn.Commit", "err", res)
		return nil, err
	}

	// ResponseCommit has no error or log, just data
	return res.Data, nil
}

func (blockExec *BlockExecutor) pruneBlocks(retainHeight int64) (uint64, error) {
	base := blockExec.blockStore.Base()
	if retainHeight <= base {
		return 0, nil
	}
	pruned, err := blockExec.blockStore.PruneBlocks(retainHeight)
	if err != nil {
		return 0, fmt.Errorf("failed to prune block store: %w", err)
	}

	err = blockExec.Store().PruneStates(retainHeight)
	if err != nil {
		return 0, fmt.Errorf("failed to prune state store: %w", err)
	}
	return pruned, nil
}

func validatePubKey(pk crypto.PubKey) error {
	v, ok := pk.(crypto.Validator)
	if !ok {
		return nil
	}
	if err := v.Validate(); err != nil {
		return err
	}
	return nil
}
