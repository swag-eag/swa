package types_test

import (
	"fmt"
	"testing"

	"github.com/swag-eag/swa/v2/app"
	"github.com/swag-eag/swa/v2/x/swa/types"
	"github.com/stretchr/testify/require"
)

func TestValidateMsgUpdateTokenMapping(t *testing.T) {
	app.SetConfig()

	testCases := []struct {
		name     string
		msg      *types.MsgUpdateTokenMapping
		expValid bool
	}{
		{
			"valid gravity denom",
			types.NewMsgUpdateTokenMapping("swac12luku6uxehhak02py4rcz65zu0swh7wjxk9sxj", "gravity0x6E7eef2b30585B2A4D45Ba9312015d5354FDB067", "0x57f96e6B86CdeFdB3d412547816a82E3E0EbF9D2", "", 0),
			true,
		},
		{
			"valid ibc denom",
			types.NewMsgUpdateTokenMapping("swac12luku6uxehhak02py4rcz65zu0swh7wjxk9sxj", "ibc/0000000000000000000000000000000000000000000000000000000000000000", "0x57f96e6B86CdeFdB3d412547816a82E3E0EbF9D2", "", 0),
			true,
		},
		{
			"invalid sender",
			types.NewMsgUpdateTokenMapping("swac12luku6uxehhak02py4r", "gravity0x6E7eef2b30585B2A4D45Ba9312015d5354FDB067", "0x57f96e6B86CdeFdB3d412547816a82E3E0EbF9D2", "", 0),
			false,
		},
		{
			"invalid denom",
			types.NewMsgUpdateTokenMapping("swac12luku6uxehhak02py4rcz65zu0swh7wjxk9sxj", "aaa", "0x57f96e6B86CdeFdB3d412547816a82E3E0EbF9D2", "", 0),
			false,
		},
		{
			"invalid contract address",
			types.NewMsgUpdateTokenMapping("swac12luku6uxehhak02py4rcz65zu0swh7wjxk9sxj", "gravity0x6E7eef2b30585B2A4D45Ba9312015d5354FDB067", "0x57f96e6B86CdeFdB3d4125", "", 0),
			false,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Case %s", tc.name), func(t1 *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expValid {
				require.NoError(t1, err)
			} else {
				require.Error(t1, err)
			}
		})
	}
}
