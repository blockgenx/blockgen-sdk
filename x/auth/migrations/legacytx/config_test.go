package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/blockgenx/blockgen-sdk/codec"
	cryptoAmino "github.com/blockgenx/blockgen-sdk/crypto/codec"
	"github.com/blockgenx/blockgen-sdk/testutil/testdata"
	sdk "github.com/blockgenx/blockgen-sdk/types"
	"github.com/blockgenx/blockgen-sdk/x/auth/migrations/legacytx"
	"github.com/blockgenx/blockgen-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
