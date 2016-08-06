package reqhandler

import (
	"fmt"
	"log"
	"net"
)

func HandleUDPConnection(udpConn net.PacketConn, l *log.Logger) {
	_, err := readRequest(udpConn)
	if err != nil {
		errMsg := fmt.Errorf("Unable to read udp connection")
		l.Println(errMsg)
	}
}

func readRequest(udpConn net.PacketConn) ([]byte, error) {
	b := make([]byte, 512)
	_, _, err := udpConn.ReadFrom(b)
	return b, err
}
