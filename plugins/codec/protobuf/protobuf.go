package protobuf

import (
	"github.com/gogo/protobuf/proto"
)

type ProtobufCodec struct {
}

// 编码器的名称
func (g ProtobufCodec) String() string {
	return "protobuf"
}

func (g ProtobufCodec) Marshal(msgObj interface{}) (data []byte, err error) {

	return proto.Marshal(msgObj.(proto.Message))

}

func (g ProtobufCodec) Unmarshal(data []byte, msgObj interface{}) error {

	return proto.Unmarshal(data, msgObj.(proto.Message))
}
