package nameserver

type mockPacker struct{}

func (mp *mockPacker) unpack(b []byte) (*message, error) {
	return &message{}, nil
}

func (mp *mockPacker) pack(m *message) ([]byte, error) {
	return []byte("hello"), nil
}
