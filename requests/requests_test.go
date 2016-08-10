package requests

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type RequestsSuite struct{}

var _ = Suite(&RequestsSuite{})

// TODO we should test setting and reading all of the headers in a single test

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
