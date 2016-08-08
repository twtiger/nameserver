package reqhandler

import (
	"fmt"
	"log"
	"net"
)

const DNSMsgLength = 512   // in bytes
const DNSHeaderLength = 12 // in bytes

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

func extractID(d []byte) uint16 {
	return (uint16(d[0]) << 8) | uint16(d[1])
}

func extractQuery(d []byte) uint16 {
	return uint16(d[3]) << (8 - HeaderFieldLengths[QUERY])
}

// TODO finish all header fields
func extractHeaders(d []byte) Headers {
	headers := Headers{}
	headers.ID = extractID(d)
	headers.Query = extractQuery(d)
	return headers
}

type Headers struct {
	ID       uint16
	Query    uint16
	Opcode   uint16
	AA       uint16
	Trunc    uint16
	RDesc    uint16
	RAvail   uint16
	RespCode uint16
	QdCount  uint16
	AnCount  uint16
	NsCount  uint16
	ArCount  uint16
}
