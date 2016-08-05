package reqhandler

import (
	"net"
)

func HandleUDPConnection(udpConn net.PacketConn) {
	_, _ = readRequest(udpConn)

}

func readRequest(udpConn net.PacketConn) ([]byte, error) {
	b := make([]byte, 512)
	_, _, err := udpConn.ReadFrom(b)
	return b, err
}
