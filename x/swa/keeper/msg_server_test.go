package keeper_test

import (
	swamodulekeeper "github.com/swag-eag/swa/v2/x/swa/keeper"
	"github.com/swag-eag/swa/v2/x/swa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (suite *KeeperTestSuite) TestUpdateParams() {
	testCases := []struct {
		name      string
		req       *types.MsgUpdateParams
		expectErr bool
		expErrMsg string
	}{
		{
			name: "gov module account address as valid authority",
			req: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					IbcCroDenom:          types.IbcCroDenomDefaultValue,
					IbcTimeout:           10,
					SwaAdmin:          sdk.AccAddress(suite.address.Bytes()).String(),
					EnableAutoDeployment: true,
				},
			},
			expectErr: false,
			expErrMsg: "",
		},
		{
			name: "set invalid authority",
			req: &types.MsgUpdateParams{
				Authority: "foo",
			},
			expectErr: true,
			expErrMsg: "invalid authority",
		},
		{
			name: "set invalid ibc cro denomination",
			req: &types.MsgUpdateParams{
				Authority: suite.app.SwaKeeper.GetAuthority(),
				Params: types.Params{
					IbcCroDenom:          "foo",
					IbcTimeout:           10,
					SwaAdmin:          sdk.AccAddress(suite.address.Bytes()).String(),
					EnableAutoDeployment: true,
				},
			},
			expectErr: true,
			expErrMsg: "invalid ibc denom",
		},
		{
			name: "set invalid swa admin address",
			req: &types.MsgUpdateParams{
				Authority: suite.app.SwaKeeper.GetAuthority(),
				Params: types.Params{
					IbcCroDenom:          types.IbcCroDenomDefaultValue,
					IbcTimeout:           10,
					SwaAdmin:          "foo",
					EnableAutoDeployment: true,
				},
			},
			expectErr: true,
			expErrMsg: "invalid bech32 string",
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			msgServer := swamodulekeeper.NewMsgServerImpl(suite.app.SwaKeeper)
			_, err := msgServer.UpdateParams(suite.ctx, tc.req)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
