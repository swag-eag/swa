package evmhandler

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	swakeeper "github.com/swag-eag/swa/v2/x/swa/keeper"
	"github.com/swag-eag/swa/v2/x/swa/types"
)

var _ types.EvmLogHandler = SendCroToIbcHandler{}

const SendCroToIbcEventName = "__SwaSendCroToIbc"

// SendCroToIbcEvent represent the signature of
// `event __SwaSendCroToIbc(string recipient, uint256 amount)`
var SendCroToIbcEvent abi.Event

func init() {
	addressType, _ := abi.NewType("address", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	stringType, _ := abi.NewType("string", "", nil)

	SendCroToIbcEvent = abi.NewEvent(
		SendCroToIbcEventName,
		SendCroToIbcEventName,
		false,
		abi.Arguments{abi.Argument{
			Name:    "sender",
			Type:    addressType,
			Indexed: false,
		}, abi.Argument{
			Name:    "recipient",
			Type:    stringType,
			Indexed: false,
		}, abi.Argument{
			Name:    "amount",
			Type:    uint256Type,
			Indexed: false,
		}},
	)
}

// SendCroToIbcHandler handles `__SwaSendCroToIbc` log
type SendCroToIbcHandler struct {
	bankKeeper   types.BankKeeper
	swaKeeper swakeeper.Keeper
}

func NewSendCroToIbcHandler(bankKeeper types.BankKeeper, swaKeeper swakeeper.Keeper) *SendCroToIbcHandler {
	return &SendCroToIbcHandler{
		bankKeeper:   bankKeeper,
		swaKeeper: swaKeeper,
	}
}

func (h SendCroToIbcHandler) EventID() common.Hash {
	return SendCroToIbcEvent.ID
}

func (h SendCroToIbcHandler) Handle(
	ctx sdk.Context,
	contract common.Address,
	topics []common.Hash,
	data []byte,
	_ func(contractAddress common.Address, logSig common.Hash, logData []byte),
) error {
	unpacked, err := SendCroToIbcEvent.Inputs.Unpack(data)
	if err != nil {
		// log and ignore
		h.swaKeeper.Logger(ctx).Error("log signature matches but failed to decode", "error", err)
		return nil
	}

	contractAddr := sdk.AccAddress(contract.Bytes())
	sender := sdk.AccAddress(unpacked[0].(common.Address).Bytes())
	recipient := unpacked[1].(string)
	amount := sdk.NewIntFromBigInt(unpacked[2].(*big.Int))
	evmDenom := h.swaKeeper.GetEvmParams(ctx).EvmDenom
	coins := sdk.NewCoins(sdk.NewCoin(evmDenom, amount))
	// First, transfer IBC coin to user so that he will be the refunded address if transfer fails
	if err = h.bankKeeper.SendCoins(ctx, contractAddr, sender, coins); err != nil {
		return err
	}
	// Initiate IBC transfer from sender account
	if err = h.swaKeeper.IbcTransferCoins(ctx, sender.String(), recipient, coins, ""); err != nil {
		return err
	}
	return nil
}
