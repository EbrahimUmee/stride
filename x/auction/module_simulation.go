package auction

import (
	"math/rand"

	"github.com/Stride-Labs/stride/v5/testutil/sample"
	auctionsimulation "github.com/Stride-Labs/stride/v5/x/auction/simulation"
	"github.com/Stride-Labs/stride/v5/x/auction/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = auctionsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
    opWeightMsgListAuctionPools = "op_weight_msg_list_auction_pools"
	// TODO: Determine the simulation weight value
	defaultWeightMsgListAuctionPools int = 100

	opWeightMsgStartAuction = "op_weight_msg_start_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStartAuction int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	auctionGenesis := types.GenesisState{
		Params:	types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&auctionGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	
	return []simtypes.ParamChange{
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgListAuctionPools int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgListAuctionPools, &weightMsgListAuctionPools, nil,
		func(_ *rand.Rand) {
			weightMsgListAuctionPools = defaultWeightMsgListAuctionPools
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgListAuctionPools,
		auctionsimulation.SimulateMsgListAuctionPools(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStartAuction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStartAuction, &weightMsgStartAuction, nil,
		func(_ *rand.Rand) {
			weightMsgStartAuction = defaultWeightMsgStartAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartAuction,
		auctionsimulation.SimulateMsgStartAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}