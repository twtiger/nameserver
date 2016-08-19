package nameserver

import (
	m "github.com/twtiger/toy-dns-nameserver/message"
)

// Packer is a dutiful type that packages and unpackages DNS messages
type Packer interface {
	Unpack(b []byte) (*m.Message, error)
	Pack(*m.Message) ([]byte, error)
}
