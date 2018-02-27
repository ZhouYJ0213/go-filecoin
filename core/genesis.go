package core

import (
	"context"
	"math/big"

	"gx/ipfs/QmZhoiN2zi5SBBBKb181dQm4QdvWAvEwbppZvKpp4gRyNY/go-hamt-ipld"

	"github.com/filecoin-project/go-filecoin/types"
)

// GenesisInitFunc is the signature for function that is used to create a genesis block.
type GenesisInitFunc func(cst *hamt.CborIpldStore) (*types.Block, error)

// an account with some initial funds in it
var testAccount = types.Address("satoshi")

// the filecoin network
var networkAccount = types.Address("filecoin")

var defaultAccounts = map[types.Address]int64{
	networkAccount: 100000,
	testAccount:    500,
}

var StorageMarketAddress = types.Address("storage")

// InitGenesis is the default function to create the genesis block.
func InitGenesis(cst *hamt.CborIpldStore) (*types.Block, error) {
	ctx := context.Background()
	st := types.NewEmptyStateTree(cst)

	for addr, val := range defaultAccounts {
		a, err := NewAccountActor(big.NewInt(val))
		if err != nil {
			return nil, err
		}

		if err := st.SetActor(ctx, addr, a); err != nil {
			return nil, err
		}
	}

	stAct, err := NewStorageMarketActor()
	if err != nil {
		return nil, err
	}
	if err := st.SetActor(ctx, StorageMarketAddress, stAct); err != nil {
		return nil, err
	}

	c, err := st.Flush(ctx)
	if err != nil {
		return nil, err
	}

	genesis := &types.Block{
		StateRoot: c,
		Nonce:     1337,
	}

	if _, err := cst.Put(ctx, genesis); err != nil {
		return nil, err
	}

	return genesis, nil
}
