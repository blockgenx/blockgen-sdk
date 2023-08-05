package codec

import (
	"bytes"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"github.com/blockgenx/blockgen-sdk/codec/types"
)

var defaultJM = &jsonpb.Marshaler{OrigName: true, EmitDefaults: true, AnyResolver: nil}

// ProtoMarshalJSON provides an auxiliary function to return Proto3 JSON encoded
// bytes of a message.
func ProtoMarshalJSON(msg proto.Message, resolver jsonpb.AnyResolver) ([]byte, error) {
	// We use the OrigName because camel casing fields just doesn't make sense.
	// EmitDefaults is also often the more expected behavior for CLI users
	jm := defaultJM
	if resolver != nil {
		jm = &jsonpb.Marshaler{OrigName: true, EmitDefaults: true, AnyResolver: resolver}
	}
	err := types.UnpackInterfaces(msg, types.ProtoJSONPacker{JSONPBMarshaler: jm})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := jm.Marshal(buf, msg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
