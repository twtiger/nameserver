package requests

import (
	"bytes"
	"encoding/binary"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type RequestsSuite struct{}

var _ = Suite(&RequestsSuite{})

func (s *RequestsSuite) TestDataFromUDPConnection(c *C) {
	udpPacket := []byte("TEST_CONNECTION")
	m := &mockUDPConn{data: udpPacket}
	output, err := readRequest(m)

	c.Assert(err, IsNil)
	c.Assert(bytes.Equal(udpPacket, output[:len(udpPacket)]), Equals, true)
}

func (s *RequestsSuite) TestDataReadUDPConnectionHasUDPPacketSize(c *C) {
	m := &mockUDPConn{}
	output, err := readRequest(m)

	c.Assert(err, IsNil)
	c.Assert(len(output), Equals, 512)
}

// this can be configurable for specific tests
func getTestHeaders() map[FieldName][]byte {
	return map[FieldName][]byte{
		ID:       []byte{0x00, 0x01},
		QUERY:    []byte{0x00, 0x01},
		OPCODE:   []byte{0x00, 0x01},
		AUTHANS:  []byte{0x00, 0x01},
		TRUNC:    []byte{0x00, 0x01},
		RDESC:    []byte{0x00, 0x01},
		RAVAIL:   []byte{0x00, 0x01},
		RESPCODE: []byte{0x00, 0x00},
		QDCOUNT:  []byte{0x00, 0x01},
		ANCOUNT:  []byte{0x00, 0x01},
		NSCOUNT:  []byte{0x00, 0x01},
		ARCOUNT:  []byte{0x00, 0x01},
	}
}

func buildTestHeaders() ([]byte, map[FieldName][]byte) {
	h := getTestHeaders()
	data := make([]byte, 12)
	data[0] = h[ID][0]
	data[1] = h[ID][1]
	data[2] = byte(uint(h[QUERY][1]) << (8 - HeaderFieldLengths[QUERY]))
	return data, h
}

func (s *RequestsSuite) TestReadIDFromUDPHeaders(c *C) {
	udpHeaders, headers := buildTestHeaders()
	output := extractHeaders(udpHeaders)

	expected := binary.BigEndian.Uint16(headers[ID])
	c.Assert(output.ID, Equals, expected)
}

func (s *RequestsSuite) TestReadQueryFromUDPHeaders(c *C) {
	udpHeaders, headers := buildTestHeaders()
	output := extractHeaders(udpHeaders)

	expected := binary.BigEndian.Uint16(headers[QUERY])
	c.Assert(output.ID, Equals, expected)
}
