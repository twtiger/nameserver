package nameserver

type mockPacker struct{}

func (mp *mockPacker) unpack(b []byte) (responder, error) {
	return &mockMsg{}, nil
}

func (mp *mockPacker) pack(m responder) ([]byte, error) {
	return []byte("hello"), nil
}

type mockMsg struct {
}

func (mm *mockMsg) respond() error {
	return nil
}
