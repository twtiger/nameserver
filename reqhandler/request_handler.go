package reqhandler

import (
	"fmt"
	"log"
	"net"
)

func HandleUDPConnection(udpConn net.PacketConn) {
	_, err := readRequest(udpConn)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to read udp connection")
		log.Printf(errMsg)
	}
}

func readRequest(udpConn net.PacketConn) ([]byte, error) {
	b := make([]byte, 512)
	_, _, err := udpConn.ReadFrom(b)
	return b, err
}
