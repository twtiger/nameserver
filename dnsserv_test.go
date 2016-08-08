package main

import (
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
	c.Assert(port, Equals, "53")
}
