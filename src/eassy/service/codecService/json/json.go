package json

type JsonCodec struct {
}

func NewCodec() *JsonCodec {
	p := new(JsonCodec)
	return p
}

// goroutine safe
func (p *JsonCodec) Unmarshal(route interface{}, data []byte) (interface{}, error) {

	return nil, nil
}

// goroutine safe
func (p *JsonCodec) Marshal(msg interface{}) ([]byte, error) {
	return nil, nil
}
