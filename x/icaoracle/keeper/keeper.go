package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	"github.com/tendermint/tendermint/libs/log"

	icacallbackskeeper "github.com/Stride-Labs/stride/v5/x/icacallbacks/keeper"
	"github.com/Stride-Labs/stride/v5/x/icaoracle/types"
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	paramstore paramtypes.Subspace

	scopedKeeper capabilitykeeper.ScopedKeeper
	ICS4Wrapper  types.ICS4Wrapper

	IBCKeeper           ibckeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICACallbacksKeeper  icacallbackskeeper.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	paramstore paramtypes.Subspace,

	scopedKeeper capabilitykeeper.ScopedKeeper,
	ics4Wrapper types.ICS4Wrapper,

	ibcKeeper ibckeeper.Keeper,
	icaControllerKeeper icacontrollerkeeper.Keeper,
	icaCallbacksKeeper icacallbackskeeper.Keeper,
) *Keeper {
	return &Keeper{
		cdc:                 cdc,
		storeKey:            key,
		paramstore:          paramstore,
		scopedKeeper:        scopedKeeper,
		ICS4Wrapper:         ics4Wrapper,
		IBCKeeper:           ibcKeeper,
		ICAControllerKeeper: icaControllerKeeper,
		ICACallbacksKeeper:  icaCallbacksKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}