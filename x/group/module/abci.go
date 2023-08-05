package module

import (
	sdk "github.com/blockgenx/blockgen-sdk/types"
	"github.com/blockgenx/blockgen-sdk/x/group/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	if err := k.TallyProposalsAtVPEnd(ctx); err != nil {
		panic(err)
	}

	if err := k.PruneProposals(ctx); err != nil {
		panic(err)
	}
}
