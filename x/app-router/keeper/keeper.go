package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"

	"github.com/Stride-Labs/stride/v4/x/app-router/types"
	stakeibckeeper "github.com/Stride-Labs/stride/v4/x/stakeibc/keeper"
)

type (
	Keeper struct {
		Cdc            codec.BinaryCodec
		storeKey       sdk.StoreKey
		paramstore     paramtypes.Subspace
		stakeibcKeeper stakeibckeeper.Keeper
		transferKeeper ibctransferkeeper.Keeper
	}
)

func NewKeeper(
	Cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ps paramtypes.Subspace,
	stakeibcKeeper stakeibckeeper.Keeper,
	transferKeeper ibctransferkeeper.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		Cdc:            Cdc,
		storeKey:       storeKey,
		paramstore:     ps,
		stakeibcKeeper: stakeibcKeeper,
		transferKeeper: transferKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}