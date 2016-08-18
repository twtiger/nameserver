package responder

import (
	m "github.com/twtiger/toy-dns-nameserver/message"
)

// Responder is a sensitive type that will respond to DNS messages
type Responder interface {
	Respond(*m.Message) (*m.Message, error)
}
