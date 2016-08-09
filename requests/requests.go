package requests

import (
	"net"
)

// DNSMsgLength is the entire DNS message in bytes
const DNSMsgLength = 512

// DNSHeaderLength is the DNS header length in bytes
const DNSHeaderLength = 12

// ReadUDP takes a udp connection
// Should eventually return a request object
func ReadUDP(udpConn net.PacketConn) ([]byte, error) {
	b := make([]byte, DNSMsgLength)
	_, _, err := udpConn.ReadFrom(b)
	return b, err
}

func extractID(d []byte) uint16 {
	return (uint16(d[0]) << 8) | uint16(d[1])
}

func extractQuery(d []byte) uint16 {
	b := uint16(d[2]) & uint16(1<<7)
	return b >> 7
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
