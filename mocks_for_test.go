package nameserver

type mockPacker struct{}

func (mp *mockPacker) deserialize(b []byte) (*message, error) {
	return &message{}, nil
}

func (mp *mockPacker) serialize(m *message) ([]byte, error) {
	return []byte("hello"), nil
}
