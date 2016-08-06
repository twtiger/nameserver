package dnsserv

import (
	"bytes"
	"github.com/twtiger/toy-dns-nameserver/reqhandler"
	"log"
	"net"
)

func getLogger() *log.Logger {
	var buf bytes.Buffer
	return log.New(&buf, "logger: ", log.Lshortfile)
}

func main() {
	l := getLogger()

	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 53,
	}

	for {
		udpConn, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			// handle error
		}
		go reqhandler.HandleUDPConnection(udpConn, l)
	}
}
