package app

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/kv"
	tmdb "github.com/tendermint/tm-db"
)

// StoreApp 上层应用
type StoreApp struct {
	types.Application
	store store.CommitKVStore
}

// NewStoreApp 创建应用对象
func NewStoreApp(db tmdb.DB) (*StoreApp, error) {
	app := &StoreApp{
		Application: types.NewBaseApplication(),
		store:       nil,
	}

	if err := app.initDB(db); err != nil {
		return nil, err
	}

	return app, nil
}

// 初始化存储
func (app *StoreApp) initDB(db tmdb.DB) error {
	var err error
	app.store, err = iavl.LoadStore(db, store.CommitID{}, false)
	if err != nil {
		return err
	}

	return err
}

func (app *StoreApp) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	key, value := []byte{}, []byte{}
	parts := bytes.Split(req.Tx, []byte(":"))
	if len(parts) == 2 {
		key, value = parts[0], parts[1]
	} else {
		key, value = req.Tx, req.Tx
	}
	app.store.Set(key, value)
	events := []types.Event{
		{
			Type: "set",
			Attributes: []kv.Pair{
				{
					Key:   []byte("key"),
					Value: key,
				},
				{
					Key:   []byte("value"),
					Value: value,
				},
			},
		},
	}

	return types.ResponseDeliverTx{
		Code:      code.CodeTypeOK,
		Log:       string(req.Tx),
		Info:      "store kv",
		Events:    events,
		Codespace: "app",
	}
}

func (app *StoreApp) Commit() types.ResponseCommit {
	commitID := app.store.Commit()
	return types.ResponseCommit{Data: commitID.Hash}
}

func (app *StoreApp) Query(req types.RequestQuery) types.ResponseQuery {
	iavlStore := app.store.(*iavl.Store)
	res := iavlStore.Query(types.RequestQuery{
		Data:  req.Data,
		Path:  "/key",
		Prove: true,
	})

	return res
}
