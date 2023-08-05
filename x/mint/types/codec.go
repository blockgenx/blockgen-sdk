package types

import (
	"github.com/blockgenx/blockgen-sdk/codec"
	cryptocodec "github.com/blockgenx/blockgen-sdk/crypto/codec"
)

var amino = codec.NewLegacyAmino()

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
