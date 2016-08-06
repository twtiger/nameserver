package dnsserv

import (
	"bytes"
	"github.com/twtiger/toy-dns-nameserver/reqhandler"
	"log"
	"net"
)

func initLogger() *log.Logger {
	var buf bytes.Buffer
	return log.New(&buf, "logger: ", log.Lshortfile)
}

func getUDPAddr() *net.UDPAddr {
	return &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 53,
	}
}

func start() error {
	l := initLogger()

	for {
		udpConn, err := net.ListenUDP("udp", getUDPAddr())
		if err != nil {
			return err
		}
		go reqhandler.HandleUDPConnection(udpConn, l)
	}
}
