package requests

// DNSHeaderLength is the DNS header length in bytes
const DNSHeaderLength = 12

func ParseRequest(b []byte) {
	// TODO
}

func extractID(d []byte) uint16 {
	return (uint16(d[0]) << 8) | uint16(d[1])
}

func extractQuery(d []byte) uint16 {
	queryMask := uint16(1 << 7)
	b := uint16(d[2]) & queryMask
	return b >> 7
}

func extractOpcode(d []byte) uint16 {
	opcodeMask := uint16(1<<4) | uint16(1<<3)
	b := uint16(d[2]) & opcodeMask
	return b >> 3
}

func extractAA(d []byte) uint16 {
	aaMask := uint16(1 << 2)
	b := uint16(d[2]) & aaMask
	return b >> 2
}

func extractHeaders(d []byte) Headers {
	headers := Headers{}
	headers.ID = extractID(d)
	headers.QR = extractQuery(d)
	headers.OPCODE = extractOpcode(d)
	headers.AA = extractAA(d)
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
