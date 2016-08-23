package nameserver

type mockMessage struct{}

func (m *mockMessage) deserialize(b []byte) error {
	return nil
}

func (m *mockMessage) serialize() ([]byte, error) {
	return []byte("hello"), nil
}

func (m *mockMessage) respond() error {
	return nil
}
