package binary

import (
	"github.com/davyxu/goobjfmt"
)

type BinaryCodec struct {
}

func (b BinaryCodec) String() string {
	return "binary"
}

func (b BinaryCodec) Marshal(msgObj interface{}) (data []byte, err error) {
	return goobjfmt.BinaryWrite(msgObj)

}

func (b BinaryCodec) Unmarshal(data []byte, msgObj interface{}) error {
	return goobjfmt.BinaryRead(data, msgObj)
}
