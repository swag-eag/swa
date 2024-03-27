package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/swag-eag/swa/v2/x/swa/exported"
	"github.com/swag-eag/swa/v2/x/swa/types"
)

// Migrate migrates the x/swa module state from the consensus version 1 to
// version 2. Specifically, it takes the parameters that are currently stored
// and managed by the x/params modules and stores them directly into the x/swa
// module state.
func Migrate(ctx sdk.Context, store sdk.KVStore, legacySubspace exported.Subspace, cdc codec.BinaryCodec) error {
	var currParams types.Params
	legacySubspace.GetParamSetIfExists(ctx, &currParams)

	if err := currParams.Validate(); err != nil {
		return err
	}
	if currParams.GetMaxCallbackGas() == 0 {
		currParams.MaxCallbackGas = types.MaxCallbackGasDefaultValue
	}
	bz := cdc.MustMarshal(&currParams)
	store.Set(types.ParamsKey, bz)

	return nil
}
