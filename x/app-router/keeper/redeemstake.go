package keeper

import (
	"github.com/armon/go-metrics"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"

	"github.com/Stride-Labs/stride/v4/x/app-router/types"
	stakeibckeeper "github.com/Stride-Labs/stride/v4/x/stakeibc/keeper"
	stakeibctypes "github.com/Stride-Labs/stride/v4/x/stakeibc/types"
)

func (k Keeper) TryRedeemStake(
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

	// In this case, we can't process a liquid staking transaction, because we're dealing IBC tokens from other chains
	if !transfertypes.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), newData.Denom) {
		return channeltypes.NewErrorAcknowledgement("the ibc tokens are not supported for redeem stake")
	}

	voucherPrefix := transfertypes.GetDenomPrefix(packet.GetSourcePort(), packet.GetSourceChannel())
	stAssetDenom := newData.Denom[len(voucherPrefix):]
	if !stakeibctypes.IsStAssetDenom(stAssetDenom) {
		return channeltypes.NewErrorAcknowledgement("not a liquid staking token")
	}

	hostZoneDenom := stakeibctypes.HostZoneDenomFromStAssetDenom(stAssetDenom)

	amount, ok := sdk.NewIntFromString(newData.Amount)
	if !ok {
		return channeltypes.NewErrorAcknowledgement("not a parsable amount field")
	}

	// Note: newData.denom is ibc denom for st assets - e.g. ibc/xxx
	var token = sdk.NewCoin(newData.Denom, amount)

	err := k.RunRedeemStake(ctx, parsedReceiver.StrideAccAddress, parsedReceiver.ResultReceiver, hostZoneDenom, token, []metrics.Label{})
	if err != nil {
		ack = channeltypes.NewErrorAcknowledgement(err.Error())
	}
	return ack
}

func (k Keeper) RunRedeemStake(ctx sdk.Context, addr sdk.AccAddress, receiver string, hostZoneDenom string, token sdk.Coin, labels []metrics.Label) error {
	hostZone, err := k.stakeibcKeeper.GetHostZoneFromHostDenom(ctx, hostZoneDenom)
	if err != nil {
		return err
	}

	msg := &stakeibctypes.MsgRedeemStake{
		Creator:  addr.String(),
		Amount:   token.Amount,
		HostZone: hostZone.ChainId,
		Receiver: receiver,
	}

	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := stakeibckeeper.NewMsgServerImpl(k.stakeibcKeeper)
	_, err = msgServer.RedeemStake(
		sdk.WrapSDKContext(ctx),
		msg,
	)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}
	return nil
}