package keeper_test

import (
	"errors"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/swag-eag/swa/v2/app"
	swamodulekeeper "github.com/swag-eag/swa/v2/x/swa/keeper"
	keepertest "github.com/swag-eag/swa/v2/x/swa/keeper/mock"
	"github.com/swag-eag/swa/v2/x/swa/types"
)

func (suite *KeeperTestSuite) TestGetSourceChannelID() {
	testCases := []struct {
		name          string
		ibcDenom      string
		expectedError error
		postCheck     func(channelID string)
	}{
		{
			"wrong ibc denom",
			"test",
			errors.New("test is invalid: ibc cro denom is invalid"),
			func(channelID string) {},
		},
		{
			"correct ibc denom",
			types.IbcCroDenomDefaultValue,
			nil,
			func(channelID string) {
				suite.Require().Equal(channelID, "channel-0")
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			// Create Swa Keeper with mock transfer keeper
			swaKeeper := *swamodulekeeper.NewKeeper(
				app.MakeEncodingConfig().Codec,
				suite.app.GetKey(types.StoreKey),
				suite.app.GetKey(types.MemStoreKey),
				suite.app.BankKeeper,
				keepertest.IbcKeeperMock{},
				suite.app.GravityKeeper,
				suite.app.EvmKeeper,
				suite.app.AccountKeeper,
				authtypes.NewModuleAddress(govtypes.ModuleName).String(),
			)
			suite.app.SwaKeeper = swaKeeper

			channelID, err := suite.app.SwaKeeper.GetSourceChannelID(suite.ctx, tc.ibcDenom)
			if tc.expectedError != nil {
				suite.Require().EqualError(err, tc.expectedError.Error())
			} else {
				suite.Require().NoError(err)
				tc.postCheck(channelID)
			}
		})
	}
}
