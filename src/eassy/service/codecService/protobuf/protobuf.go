package protobuf

import (
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"hash/crc32"
	"reflect"
)

type ProtobufCodec struct {
	littleEndian bool
}

func NewCodec() *ProtobufCodec {
	p := new(ProtobufCodec)
	p.littleEndian = false
	return p
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *ProtobufCodec) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

// goroutine safe
func (p *ProtobufCodec) Unmarshal(msgReq interface{}, data []byte) (interface{}, error) {
	proto.UnmarshalMerge(data, msgReq.(proto.Message))

	return nil, nil
}

// goroutine safe
func (p *ProtobufCodec) Marshal(msg interface{}) ([]byte, error) {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return nil, errors.New("pb message pointer required")
	}
	msgID := msgType.Elem().Name()
	_id := crc32.ChecksumIEEE([]byte(msgID))

	id := make([]byte, 4)
	if p.littleEndian {
		binary.LittleEndian.PutUint32(id, _id)
	} else {
		binary.BigEndian.PutUint32(id, _id)
	}

	// data
	data, err := proto.Marshal(msg.(proto.Message))
	return data, err
}