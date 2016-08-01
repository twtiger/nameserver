package dnsserv

import (
	"github.com/twtiger/toy-dns-nameserver/reqhandler"
	"net"
)

func main() {
	ln, err := net.Listen("udp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go reqhandler.HandleConnection(conn)
	}
}
