package nameserver

type mockPacker struct{}

func (mp *mockPacker) unpack(b []byte) (Responder, error) {
	return &mockMsg{}, nil
}

func (mp *mockPacker) pack(m Responder) ([]byte, error) {
	return []byte("hello"), nil
}

type mockMsg struct {
}

func (mm *mockMsg) Respond() error {
	return nil
}
