package simulation

import (
	"math/rand"

	"github.com/Stride-Labs/stride/v5/x/auction/keeper"
	"github.com/Stride-Labs/stride/v5/x/auction/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgStartAuction(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgStartAuction{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the StartAuction simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "StartAuction simulation not implemented"), nil, nil
	}
}