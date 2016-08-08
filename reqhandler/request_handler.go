package reqhandler

import (
	"fmt"
	"log"
	"net"
)

const DNSMsgLength = 512

func HandleUDPConnection(udpConn net.PacketConn) {
	_, err := readRequest(udpConn)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to read udp connection, error: %s", err.Error())
		log.Printf(errMsg)
	}
}

func readRequest(udpConn net.PacketConn) ([]byte, error) {
	b := make([]byte, DNSMsgLength)
	_, _, err := udpConn.ReadFrom(b)
	return b, err
}
