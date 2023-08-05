package keeper_test

import (
	"github.com/blockgenx/blockgen-sdk/codec"
	"github.com/blockgenx/blockgen-sdk/simapp"
	storetypes "github.com/blockgenx/blockgen-sdk/store/types"
	"github.com/blockgenx/blockgen-sdk/testutil"
	sdk "github.com/blockgenx/blockgen-sdk/types"
	paramskeeper "github.com/blockgenx/blockgen-sdk/x/params/keeper"
)

func testComponents() (*codec.LegacyAmino, sdk.Context, storetypes.StoreKey, storetypes.StoreKey, paramskeeper.Keeper) {
	marshaler := simapp.MakeTestEncodingConfig().Codec
	legacyAmino := createTestCodec()
	mkey := sdk.NewKVStoreKey("test")
	tkey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(mkey, tkey)
	keeper := paramskeeper.NewKeeper(marshaler, legacyAmino, mkey, tkey)

	return legacyAmino, ctx, mkey, tkey, keeper
}

type invalid struct{}

type s struct {
	I int
}

func createTestCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cdc.RegisterConcrete(s{}, "test/s", nil)
	cdc.RegisterConcrete(invalid{}, "test/invalid", nil)
	return cdc
}
