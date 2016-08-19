package nameserver

import (
	mes "github.com/twtiger/toy-dns-nameserver/message"
)

type packer interface {
	unpack(b []byte) (mes.Responder, error)
	pack(mes.Responder) ([]byte, error)
}

type msgPacker struct{}

func (mp *msgPacker) unpack(b []byte) (mes.Responder, error) {
	return nil, nil
}

func (mp *msgPacker) pack(m mes.Responder) ([]byte, error) {
	return nil, nil
}
