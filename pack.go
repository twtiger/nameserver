package nameserver

type packer interface {
	unpack(b []byte) (Responder, error)
	pack(Responder) ([]byte, error)
}

type msgPacker struct{}

func (mp *msgPacker) unpack(b []byte) (Responder, error) {
	return nil, nil
}

func (mp *msgPacker) pack(m Responder) ([]byte, error) {
	return nil, nil
}
