package nameserver

import (
	"net"

	. "gopkg.in/check.v1"
)

type NameserverSuite struct{}

var _ = Suite(&NameserverSuite{})

func (s *NameserverSuite) TearDownTest(c *C) {
	if ns.ucon != nil {
		ns.teardown()
	}
}

func (s *NameserverSuite) TestSuccessfulConnectionHasValidAddress(c *C) {
	ns = localServer(true)

	err := ns.Connect()

	c.Assert(err, Equals, nil)
	c.Assert(ns.ucon.LocalAddr().Network(), Equals, "udp")
	c.Assert(ns.ucon.LocalAddr().String(), Equals, "127.0.0.1:8899")
}

func (s *NameserverSuite) TestUnsuccessfulConnectionReturnsError(c *C) {
	ns = localServer(false)

	err := ns.Connect()

	c.Assert(err, ErrorMatches, "unable to connect:.*")
}

func (s *NameserverSuite) TestCannotUseServeWithoutConnecting(c *C) {
	ns = localServer(true)

	err := ns.Serve()

	c.Assert(err, ErrorMatches, "not connected: must successfully connect with nameserver.Connect first")
}

func (s *NameserverSuite) TestThatServerIsReplyingOnListeningPort(c *C) {
	const resPort = 8845
	errChan := make(chan error)
	addrChan := make(chan *net.UDPAddr)

	ns = localServer(true)
	ns.Connect()

	resolverListener, _ := net.ListenUDP("udp", localhost(resPort))
	go func() {
		b := make([]byte, 512)
		_, ra, err := resolverListener.ReadFromUDP(b)
		defer resolverListener.Close()

		addrChan <- ra
		errChan <- err
	}()

	ns.reply([]byte("hi"), localhost(resPort))

	retAddr := <-addrChan
	errRead := <-errChan
	c.Assert(retAddr.IP.String(), Equals, "127.0.0.1")
	c.Assert(retAddr.Port, Equals, nsPort)
	c.Assert(errRead, IsNil)
}
