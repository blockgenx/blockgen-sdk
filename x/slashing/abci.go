package slashing

import (
	"time"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/blockgenx/blockgen-sdk/telemetry"
	sdk "github.com/blockgenx/blockgen-sdk/types"
	"github.com/blockgenx/blockgen-sdk/x/slashing/keeper"
	"github.com/blockgenx/blockgen-sdk/x/slashing/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// Iterate over all the validators which *should* have signed this block
	// store whether or not they have actually signed it and slash/unbond any
	// which have missed too many blocks in a row (downtime slashing)
	for _, voteInfo := range req.LastCommitInfo.GetVotes() {
		k.HandleValidatorSignature(ctx, voteInfo.Validator.Address, voteInfo.Validator.Power, voteInfo.SignedLastBlock)
	}
}
