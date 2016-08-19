package nameserver

type packer interface {
	unpack(b []byte) (responder, error)
	pack(responder) ([]byte, error)
}

type msgPacker struct{}

func (mp *msgPacker) unpack(b []byte) (responder, error) {
	return nil, nil
}

func (mp *msgPacker) pack(m responder) ([]byte, error) {
	return nil, nil
}
