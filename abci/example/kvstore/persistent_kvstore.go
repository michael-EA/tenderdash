package kvstore

import (
	"context"

	dbm "github.com/tendermint/tm-db"

	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
)

//-----------------------------------------

var _ types.Application = (*PersistentKVStoreApplication)(nil)

type PersistentKVStoreApplication struct {
	*Application
}

func NewPersistentKVStoreApplication(logger log.Logger, dbDir string) *PersistentKVStoreApplication {
	db, err := dbm.NewGoLevelDB("kvstore", dbDir)
	if err != nil {
		panic(err)
	}

	return &PersistentKVStoreApplication{
		Application: &Application{
			state:          loadState(db),
			logger:         logger,
			valsIndex:      make(map[string]*types.ValidatorUpdate),
			valUpdatesRepo: &repository{db},
		},
	}
}

func (app *PersistentKVStoreApplication) OfferSnapshot(_ context.Context, req *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	return &types.ResponseOfferSnapshot{Result: types.ResponseOfferSnapshot_ABORT}, nil
}

func (app *PersistentKVStoreApplication) ApplySnapshotChunk(_ context.Context, req *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	return &types.ResponseApplySnapshotChunk{Result: types.ResponseApplySnapshotChunk_ABORT}, nil
}
