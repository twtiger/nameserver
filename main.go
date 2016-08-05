package dnsserv

import (
	"github.com/twtiger/toy-dns-nameserver/reqhandler"
	"net"
)

func main() {

	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 53,
	}

	for {
		udpConn, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			// handle error
		}
		go reqhandler.HandleUDPConnection(udpConn)
	}
}
