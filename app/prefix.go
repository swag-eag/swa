package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	cmdcfg "github.com/swag-eag/swa/v2/cmd/swad/config"
	ethcfg "github.com/evmos/ethermint/cmd/config"
)

func SetConfig() {
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	ethcfg.SetBip44CoinType(config)
	// Make sure address is compatible with ethereum
	config.SetAddressVerifier(VerifyAddressFormat)
	config.Seal()
}
