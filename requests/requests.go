package requests

import (
	"fmt"
	"log"
	"net"
)

// DNSMsgLength is the entire DNS message in bytes
const DNSMsgLength = 512

// DNSHeaderLength is the DNS header length in bytes
const DNSHeaderLength = 12

// HandleUDPConnection takes a udp connection, handles any errors, and should return a request
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
	return uint16(d[3]) << (8 - HeaderFieldLengths[QR])
}

// TODO finish all header fields
func extractHeaders(d []byte) Headers {
	headers := Headers{}
	headers.ID = extractID(d)
	headers.QR = extractQuery(d)
	return headers
}

// Headers for DNS
type Headers struct {
	ID      uint16
	QR      uint16
	OPCODE  uint16
	AA      uint16
	TR      uint16
	RD      uint16
	RA      uint16
	Z       uint16
	RCODE   uint16
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}
