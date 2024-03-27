package types

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var (
	Ten       = big.NewInt(10)
	TenPowTen = Ten.Exp(Ten, Ten, nil)
)

const (
	ibcDenomPrefix     = "ibc/"
	ibcDenomLen        = len(ibcDenomPrefix) + 64
	gravityDenomPrefix = "gravity0x"
	gravityDenomLen    = len(gravityDenomPrefix) + 40
	swaDenomPrefix  = "swa0x"
	swaDenomLen     = len(swaDenomPrefix) + 40
)

// IsValidIBCDenom returns true if denom is a valid ibc denom
func IsValidIBCDenom(denom string) bool {
	return len(denom) == ibcDenomLen && strings.HasPrefix(denom, ibcDenomPrefix)
}

// IsValidGravityDenom returns true if denom is a valid gravity denom
func IsValidGravityDenom(denom string) bool {
	return len(denom) == gravityDenomLen && strings.HasPrefix(denom, gravityDenomPrefix)
}

// IsValidSwaDenom returns true if denom is a valid swa denom
func IsValidSwaDenom(denom string) bool {
	return len(denom) == swaDenomLen && strings.HasPrefix(denom, swaDenomPrefix)
}

// IsSourceCoin returns true if denom is a coin originated from swa
func IsSourceCoin(denom string) bool {
	return IsValidSwaDenom(denom)
}

// IsValidCoinDenom returns true if it's ok it is a valid coin denom
func IsValidCoinDenom(denom string) bool {
	return IsValidIBCDenom(denom) || IsValidGravityDenom(denom) || IsValidSwaDenom(denom)
}

// GetContractAddressFromDenom get the contract address from the coin denom
func GetContractAddressFromDenom(denom string) (string, error) {
	contractAddress := ""
	if strings.HasPrefix(denom, gravityDenomPrefix) {
		contractAddress = denom[7:]
	} else if strings.HasPrefix(denom, swaDenomPrefix) {
		contractAddress = denom[6:]
	}
	if !common.IsHexAddress(contractAddress) {
		return "", fmt.Errorf("invalid contract address (%s)", contractAddress)
	}
	return contractAddress, nil
}
