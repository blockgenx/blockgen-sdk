package std

import (
	"github.com/blockgenx/blockgen-sdk/codec"
	"github.com/blockgenx/blockgen-sdk/codec/types"
	cryptocodec "github.com/blockgenx/blockgen-sdk/crypto/codec"
	sdk "github.com/blockgenx/blockgen-sdk/types"
	txtypes "github.com/blockgenx/blockgen-sdk/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
