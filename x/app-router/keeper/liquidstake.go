package keeper

import (
	"github.com/armon/go-metrics"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"

	"github.com/Stride-Labs/stride/v4/x/app-router/types"
	stakeibckeeper "github.com/Stride-Labs/stride/v4/x/stakeibc/keeper"
	stakeibctypes "github.com/Stride-Labs/stride/v4/x/stakeibc/types"
)

func (k Keeper) TryLiquidStaking(
	ctx sdk.Context,
	packet channeltypes.Packet,
	newData transfertypes.FungibleTokenPacketData,
	parsedReceiver *types.ParsedReceiver,
	ack ibcexported.Acknowledgement,
) ibcexported.Acknowledgement {
	params := k.GetParams(ctx)
	if !params.Active {
		return channeltypes.NewErrorAcknowledgement("packet forwarding param is not active")
	}

	// In this case, we can't process a liquid staking transaction, because we're dealing with STRD tokens
	if transfertypes.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), newData.Denom) {
		return channeltypes.NewErrorAcknowledgement("the native token is not supported for liquid staking")
	}

	amount, ok := sdk.NewIntFromString(newData.Amount)
	if !ok {
		return channeltypes.NewErrorAcknowledgement("not a parsable amount field")
	}

	// Note: newData.denom is base denom e.g. uatom - not ibc/xxx
	var token = sdk.NewCoin(newData.Denom, amount)

	prefixedDenom := transfertypes.GetPrefixedDenom(packet.GetDestPort(), packet.GetDestChannel(), newData.Denom)
	ibcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()

	hostZone, err := k.stakeibcKeeper.GetHostZoneFromHostDenom(ctx, token.Denom)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err.Error())
	}

	if hostZone.IbcDenom != ibcDenom {
		return channeltypes.NewErrorAcknowledgement("ibc denom is not equal to host zone ibc denom")
	}

	err = k.RunLiquidStake(ctx, parsedReceiver.StrideAccAddress, parsedReceiver.ResultReceiver, token, []metrics.Label{})
	if err != nil {
		ack = channeltypes.NewErrorAcknowledgement(err.Error())
	}
	return ack
}

func (k Keeper) RunLiquidStake(ctx sdk.Context, addr sdk.AccAddress, ibcReceiver string, token sdk.Coin, labels []metrics.Label) error {
	msg := &stakeibctypes.MsgLiquidStake{
		Creator:   addr.String(),
		Amount:    token.Amount,
		HostDenom: token.Denom,
	}

	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := stakeibckeeper.NewMsgServerImpl(k.stakeibcKeeper)
	lsRes, err := msgServer.LiquidStake(
		sdk.WrapSDKContext(ctx),
		msg,
	)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	if ibcReceiver == "" {
		return nil
	}

	hostZone, err := k.stakeibcKeeper.GetHostZoneFromHostDenom(ctx, token.Denom)
	if err != nil {
		return err
	}

	return k.IBCTransferStAsset(ctx, lsRes.StAsset, hostZone.TransferChannelId, addr.String(), ibcReceiver)
}

func (k Keeper) IBCTransferStAsset(ctx sdk.Context, stAsset sdk.Coin, channelId string, addr string, ibcReceiver string) error {
	ibcTransferTimeoutNanos := k.stakeibcKeeper.GetParam(ctx, stakeibctypes.KeyIBCTransferTimeoutNanos)
	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + ibcTransferTimeoutNanos
	transferMsg := &transfertypes.MsgTransfer{
		SourcePort:       transfertypes.PortID,
		SourceChannel:    channelId,
		Token:            stAsset,
		Sender:           addr,
		Receiver:         ibcReceiver,
		TimeoutHeight:    clienttypes.Height{},
		TimeoutTimestamp: timeoutTimestamp,
		Memo:             "stTokenIBCTransfer",
	}

	_, err := k.transferKeeper.Transfer(sdk.WrapSDKContext(ctx), transferMsg)
	return err
}