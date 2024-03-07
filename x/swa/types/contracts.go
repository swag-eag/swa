package types

import (
	// embed compiled smart contract
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// ByteString is a byte array that serializes to hex
type ByteString []byte

// MarshalJSON serializes ByteArray to hex
func (s ByteString) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(fmt.Sprintf("%x", string(s)))
	return bytes, err
}

// UnmarshalJSON deserializes ByteArray to hex
func (s *ByteString) UnmarshalJSON(data []byte) error {
	var x string
	err := json.Unmarshal(data, &x)
	if err == nil {
		str, e := hex.DecodeString(x)
		*s = str
		err = e
	}

	return err
}

// CompiledContract contains compiled bytecode and abi
type CompiledContract struct {
	ABI abi.ABI
	Bin ByteString
}

const EVMModuleName = "cronos-evm"

var (
	//go:embed contracts/ModuleSWAC20.json
	cronosSWAC20JSON []byte

	// ModuleSWAC20Contract is the compiled cronos swac20 contract
	ModuleSWAC20Contract CompiledContract

	//go:embed contracts/ModuleSWAC21.json
	cronosSWAC21JSON []byte

	// ModuleSWAC21Contract is the compiled cronos swac21 contract
	ModuleSWAC21Contract CompiledContract

	// EVMModuleAddress is the native module address for EVM
	EVMModuleAddress common.Address
)

func init() {
	EVMModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(EVMModuleName).Bytes())

	err := json.Unmarshal(cronosSWAC20JSON, &ModuleSWAC20Contract)
	if err != nil {
		panic(err)
	}

	if len(ModuleSWAC20Contract.Bin) == 0 {
		panic("load contract failed")
	}

	err = json.Unmarshal(cronosSWAC21JSON, &ModuleSWAC21Contract)
	if err != nil {
		panic(err)
	}

	if len(ModuleSWAC21Contract.Bin) == 0 {
		panic("load contract failed")
	}
}
