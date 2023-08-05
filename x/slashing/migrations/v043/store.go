package v043

import (
	storetypes "github.com/blockgenx/blockgen-sdk/store/types"
	sdk "github.com/blockgenx/blockgen-sdk/types"
	v043distribution "github.com/blockgenx/blockgen-sdk/x/distribution/migrations/v043"
	v042slashing "github.com/blockgenx/blockgen-sdk/x/slashing/migrations/v042"
)

// MigrateStore performs in-place store migrations from v0.40 to v0.43. The
// migration includes:
//
// - Change addresses to be length-prefixed.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey) error {
	store := ctx.KVStore(storeKey)
	v043distribution.MigratePrefixAddress(store, v042slashing.ValidatorSigningInfoKeyPrefix)
	v043distribution.MigratePrefixAddressBytes(store, v042slashing.ValidatorMissedBlockBitArrayKeyPrefix)
	v043distribution.MigratePrefixAddress(store, v042slashing.AddrPubkeyRelationKeyPrefix)

	return nil
}
