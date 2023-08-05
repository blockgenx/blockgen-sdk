package module

import (
	sdk "github.com/blockgenx/blockgen-sdk/types"
	"github.com/blockgenx/blockgen-sdk/x/feegrant/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.RemoveExpiredAllowances(ctx)
}
