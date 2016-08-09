package main

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type DNSNameserverSuite struct{}

var _ = Suite(&DNSNameserverSuite{})

func (s *DNSNameserverSuite) TestUDPConnection(c *C) {
	output := getUDPAddr()
	host, port, err := net.SplitHostPort(output.String())

	c.Assert(err, IsNil)
	c.Assert(host, Equals, "127.0.0.1")
	c.Assert(port, Equals, "8888")
}

func (s *DNSNameserverSuite) TestDataFromUDPConnection(c *C) {
	udpPacket := []byte("TEST_CONNECTION")
	m := &mockUDPConn{data: udpPacket}
	output, err := readUDPPacket(m)

	c.Assert(err, IsNil)
	c.Assert(bytes.Equal(udpPacket, output[:len(udpPacket)]), Equals, true)
}

func (s *DNSNameserverSuite) TestDataReadUDPConnectionHasUDPPacketSize(c *C) {
	m := &mockUDPConn{}
	output, err := readUDPPacket(m)

	c.Assert(err, IsNil)
	c.Assert(output, HasLen, 512)
}
