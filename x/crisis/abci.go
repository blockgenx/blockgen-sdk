package crisis

import (
	"time"

	"github.com/blockgenx/blockgen-sdk/telemetry"
	sdk "github.com/blockgenx/blockgen-sdk/types"
	"github.com/blockgenx/blockgen-sdk/x/crisis/keeper"
	"github.com/blockgenx/blockgen-sdk/x/crisis/types"
)

// check all registered invariants
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	if k.InvCheckPeriod() == 0 || ctx.BlockHeight()%int64(k.InvCheckPeriod()) != 0 {
		// skip running the invariant check
		return
	}
	k.AssertInvariants(ctx)
}
