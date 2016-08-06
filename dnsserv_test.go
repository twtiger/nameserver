package dnsserv

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
	c.Assert("127.0.0.1", Equals, host)
	c.Assert("53", Equals, port)
}
