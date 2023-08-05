package codec

import (
	"github.com/blockgenx/blockgen-sdk/codec"
	cryptocodec "github.com/blockgenx/blockgen-sdk/crypto/codec"
	sdk "github.com/blockgenx/blockgen-sdk/types"
)

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	cryptocodec.RegisterCrypto(Amino)
	codec.RegisterEvidences(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
