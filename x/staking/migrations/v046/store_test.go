package v046_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/blockgenx/blockgen-sdk/simapp"
	"github.com/blockgenx/blockgen-sdk/testutil"
	sdk "github.com/blockgenx/blockgen-sdk/types"
	paramtypes "github.com/blockgenx/blockgen-sdk/x/params/types"
	v046staking "github.com/blockgenx/blockgen-sdk/x/staking/migrations/v046"
	"github.com/blockgenx/blockgen-sdk/x/staking/types"
)

func TestStoreMigration(t *testing.T) {
	encCfg := simapp.MakeTestEncodingConfig()
	stakingKey := sdk.NewKVStoreKey("staking")
	tStakingKey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(stakingKey, tStakingKey)
	paramstore := paramtypes.NewSubspace(encCfg.Codec, encCfg.Amino, stakingKey, tStakingKey, "staking")

	// Check no params
	require.False(t, paramstore.Has(ctx, types.KeyMinCommissionRate))

	// Run migrations.
	err := v046staking.MigrateStore(ctx, stakingKey, encCfg.Codec, paramstore)
	require.NoError(t, err)

	// Make sure the new params are set.
	require.True(t, paramstore.Has(ctx, types.KeyMinCommissionRate))
}
