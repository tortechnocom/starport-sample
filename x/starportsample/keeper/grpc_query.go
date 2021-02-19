package keeper

import (
	"github.com/tortechnocom/starport-sample/x/starportsample/types"
)

var _ types.QueryServer = Keeper{}
