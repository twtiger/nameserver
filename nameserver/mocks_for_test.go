package nameserver

import (
	msg "github.com/twtiger/toy-dns-nameserver/message"
)

type mockPacker struct{}

func (mp *mockPacker) unpack(b []byte) (msg.Responder, error) {
	return &mockMsg{}, nil
}

func (mp *mockPacker) pack(m msg.Responder) ([]byte, error) {
	return []byte("hello"), nil
}

type mockMsg struct {
}

func (mm *mockMsg) Respond() error {
	return nil
}
