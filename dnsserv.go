package main

import (
	"github.com/twtiger/toy-dns-nameserver/reqhandler"
	"net"
)

func getUDPAddr() *net.UDPAddr {
	return &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 53,
	}
}

func run() error {
	for {
		udpConn, err := net.ListenUDP("udp", getUDPAddr())
		if err != nil {
			return err
		}
		go reqhandler.HandleUDPConnection(udpConn)
		udpConn.Close()
	}
}
