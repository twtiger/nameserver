package requests

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type RequestsSuite struct{}

var _ = Suite(&RequestsSuite{})

func getTestHeaders() map[FieldName][]byte {
	return map[FieldName][]byte{
		ID:      []byte{0x00, 0x01},
		QR:      []byte{0x00, 0x01},
		OPCODE:  []byte{0x00, 0x01},
		AA:      []byte{0x00, 0x01},
		TC:      []byte{0x00, 0x01},
		RD:      []byte{0x00, 0x01},
		RA:      []byte{0x00, 0x01},
		Z:       []byte{0x00, 0x00},
		RCODE:   []byte{0x00, 0x00},
		QDCOUNT: []byte{0x00, 0x01},
		ANCOUNT: []byte{0x00, 0x01},
		NSCOUNT: []byte{0x00, 0x01},
		ARCOUNT: []byte{0x00, 0x01},
	}
}

func buildTestHeaders() ([]byte, map[FieldName][]byte) {
	h := getTestHeaders()
	data := make([]byte, 12)
	data[0] = h[ID][0]
	data[1] = h[ID][1]
	data[2] = byte(uint16(h[QR][1]) << uint16(7))
	return data, h
}

func (s *RequestsSuite) TestReadIDFromUDPHeaders(c *C) {
	data := make([]byte, 12)
	field := HeaderFields[ID]
	data[field.position+1] = byte(uint16(1) << uint16(field.offset))
	output := extractHeaders(data)

	c.Assert(output.ID, Equals, uint16(1))
}

func (s *RequestsSuite) TestReadQueryFromUDPHeaders(c *C) {
	data := make([]byte, 12)
	field := HeaderFields[QR]
	lastBitPos := 7 - field.offset
	data[field.position] = byte(uint16(1) << uint16(lastBitPos))
	output := extractHeaders(data)

	c.Assert(output.QR, Equals, uint16(1))
}

func (s *RequestsSuite) TestReadOpcodeFromUDPHeaders(c *C) {
	data := make([]byte, 12)
	field := HeaderFields[OPCODE]
	lastBitPos := 7 - field.offset - field.length
	data[field.position] = byte(uint16(1) << uint16(lastBitPos))
	output := extractHeaders(data)

	c.Assert(output.OPCODE, Equals, uint16(1))
}

func (s *RequestsSuite) TestReadAAFromUDPHeaders(c *C) {
	data := make([]byte, 12)
	field := HeaderFields[AA]
	lastBitPos := 7 - field.offset
	data[field.position] = byte(uint16(1) << uint16(lastBitPos))
	output := extractHeaders(data)

	c.Assert(output.AA, Equals, uint16(1))
}
