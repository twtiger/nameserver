package main

import (
	"fmt"
	"log"
	"net"
)

const dnsMsgLength = 512

func getUDPAddr() *net.UDPAddr {
	return &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8888,
	}
}

func readUDPPacket(udpConn net.PacketConn) ([]byte, error) {
	b := make([]byte, dnsMsgLength)
	_, _, err := udpConn.ReadFrom(b)
	if err != nil {
		e := fmt.Sprintf("Error in reading message %s", err.Error())
		log.Printf(e)
	}
	return b, err
}

func handleRequest(p []byte) {
	// TODO
}

func run() error {
	udpConn, err := net.ListenUDP("udp", getUDPAddr())

	if err != nil {
		return err
	}

	for {
		b, e := readUDPPacket(udpConn)
		if e != nil {
			go handleRequest(b)
		}
	}
}
