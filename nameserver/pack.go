package nameserver

import (
	m "github.com/twtiger/toy-dns-nameserver/message"
)

// Packer is able to turn bytes into a DNS message
type Packer interface {
	Unpack(b []byte) (*m.Message, error)
	Pack(*m.Message) ([]byte, error)
}
