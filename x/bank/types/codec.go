package types

import (
	"github.com/blockgenx/blockgen-sdk/codec"
	"github.com/blockgenx/blockgen-sdk/codec/legacy"
	"github.com/blockgenx/blockgen-sdk/codec/types"
	cryptocodec "github.com/blockgenx/blockgen-sdk/crypto/codec"
	sdk "github.com/blockgenx/blockgen-sdk/types"
	"github.com/blockgenx/blockgen-sdk/types/msgservice"
	"github.com/blockgenx/blockgen-sdk/x/authz"
	authzcodec "github.com/blockgenx/blockgen-sdk/x/authz/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/bank interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgSend{}, "cosmos-sdk/MsgSend")
	legacy.RegisterAminoMsg(cdc, &MsgMultiSend{}, "cosmos-sdk/MsgMultiSend")
	cdc.RegisterConcrete(&SendAuthorization{}, "cosmos-sdk/SendAuthorization", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
	)
	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&SendAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
