package main

import (
	"fmt"
	"github.com/twtiger/toy-dns-nameserver/requests"
	"log"
	"net"
)

func getUDPAddr() *net.UDPAddr {
	return &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 53,
	}
}

// HandleUDPConnection takes a udp connection, handles any errors, and should return a request
func handleUDPConnection(udpConn net.PacketConn) {
	_, err := requests.ReadUDP(udpConn)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to read udp connection, error: %s", err.Error())
		log.Printf(errMsg)
	}
}

func run() error {
	for {
		udpConn, err := net.ListenUDP("udp", getUDPAddr())
		if err != nil {
			return err
		}
		go handleUDPConnection(udpConn)
		udpConn.Close()
	}
}
