package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/swag-eag/swa/v2/x/swa/simulation"
	"github.com/swag-eag/swa/v2/x/swa/types"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	registry := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: math.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var swaGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &swaGenesis)

	require.Equal(t, "ibc/7939cb6694d2c422acd208a0072939487f6999eb9d18a44784045d87f3c67cf2", swaGenesis.Params.GetIbcCroDenom())
	require.Equal(t, uint64(0x68255aaf95e94627), swaGenesis.Params.GetIbcTimeout())
	require.Equal(t, "cosmos1tnh2q55v8wyygtt9srz5safamzdengsnqeycj3", swaGenesis.Params.GetSwaAdmin())
	require.Equal(t, true, swaGenesis.Params.GetEnableAutoDeployment())

	require.Equal(t, len(swaGenesis.ExternalContracts), 0)
	require.Equal(t, len(swaGenesis.AutoContracts), 0)
}
