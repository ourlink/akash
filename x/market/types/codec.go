package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var cdc = codec.New()

func init() {
	RegisterCodec(cdc)
}

// RegisterCodec register concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateBid{}, ModuleName+"/"+msgTypeCreateBid, nil)
	cdc.RegisterConcrete(MsgCloseBid{}, ModuleName+"/"+msgTypeCloseBid, nil)
	cdc.RegisterConcrete(MsgCloseOrder{}, ModuleName+"/"+msgTypeCloseOrder, nil)
}

// MustMarshalJSON panics if an error occurs. Besides that it behaves exactly like MarshalJSON
// i.e., encodes json to byte array
func MustMarshalJSON(o interface{}) []byte {
	return cdc.MustMarshalJSON(o)
}

// UnmarshalJSON decodes bytes into json
func UnmarshalJSON(bz []byte, ptr interface{}) error {
	return cdc.UnmarshalJSON(bz, ptr)
}

// MustUnmarshalJSON panics if an error occurs. Besides that it behaves exactly like UnmarshalJSON.
func MustUnmarshalJSON(bz []byte, ptr interface{}) {
	cdc.MustUnmarshalJSON(bz, ptr)
}
