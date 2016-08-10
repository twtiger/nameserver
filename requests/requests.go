package requests

// DNSHeaderLength is the DNS header length in bytes
const DNSHeaderLength = 12

func ParseRequest(b []byte) {
	// TODO
}

// TODO make for variable number of bytes
func extractTwoBytes(d []byte, f Field) uint16 {
	return (uint16(d[f.position]) << 8) | uint16(d[f.position+1])
}

func extractBit(d []byte, f Field) uint16 {
	pos := 7 - f.offset
	mask := uint16(1 << pos)
	b := uint16(d[f.position]) & mask
	return b >> pos
}

// This maybe could be combined with extractBit
func extractMultipleBits(d []byte, f Field) uint16 {
	b := uint16(0)
	pos := 7 - f.offset - f.length
	for n := f.length; n > 0; n-- {
		mask := uint16(1 << n)
		b = b | uint16(d[f.position])&mask
	}
	return b >> pos
}

func extractHeaders(d []byte) Headers {
	headers := Headers{}
	headers.ID = extractTwoBytes(d, HeaderFields[ID])
	headers.QR = extractBit(d, HeaderFields[QR])
	headers.OPCODE = extractMultipleBits(d, HeaderFields[OPCODE])
	headers.AA = extractBit(d, HeaderFields[AA])
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
