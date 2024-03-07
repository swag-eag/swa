package keeper

import (
	"github.com/swag-eag/swa/v2/x/icaauth/types"
)

var _ types.QueryServer = Keeper{}
