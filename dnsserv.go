package main

import (
	"github.com/twtiger/toy-dns-nameserver/requests"
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
		go requests.HandleUDPConnection(udpConn)
		udpConn.Close()
	}
}
