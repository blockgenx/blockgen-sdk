package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/blockgenx/blockgen-sdk/codec"
	codectypes "github.com/blockgenx/blockgen-sdk/codec/types"
	"github.com/blockgenx/blockgen-sdk/std"
	"github.com/blockgenx/blockgen-sdk/testutil/testdata"
	sdk "github.com/blockgenx/blockgen-sdk/types"
	"github.com/blockgenx/blockgen-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
