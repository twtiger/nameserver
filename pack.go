package nameserver

type packer interface {
	unpack(b []byte) (*message, error)
	pack(*message) ([]byte, error)
}

type msgPacker struct{}

func (mp *msgPacker) unpack(b []byte) (*message, error) {
	return nil, nil
}

func (mp *msgPacker) pack(m *message) ([]byte, error) {
	return nil, nil
}
