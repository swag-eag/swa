package swa

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/swag-eag/swa/v2/x/swa/keeper"
	"github.com/swag-eag/swa/v2/x/swa/types"
)

// NewTokenMappingChangeProposalHandler creates a new governance Handler for a TokenMappingChangeProposal
func NewTokenMappingChangeProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.TokenMappingChangeProposal:
			// check first that the denom is one of the denom supported by swa
			if !types.IsValidCoinDenom(c.Denom) {
				return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid coin denom %s", c.Denom)
			}

			msg := types.MsgUpdateTokenMapping{
				Denom:    c.Denom,
				Contract: c.Contract,
				Symbol:   c.Symbol,
				Decimal:  c.Decimal,
			}
			return k.RegisterOrUpdateTokenMapping(ctx, &msg)
		default:
			return errors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized swa proposal content type: %T", c)
		}
	}
}
