package nameserver

import (
	"net"
	"testing"

	dnsr "github.com/bogdanovich/dns_resolver"
	. "gopkg.in/check.v1"
)

var _ = dnsr.New // to write functional tests later

func Test(t *testing.T) { TestingT(t) }

type NameserverSuite struct{}

var _ = Suite(&NameserverSuite{})

var ns Nameserver

func localhost(port int) *net.UDPAddr {
	return &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: port,
	}
}

func localServer(valid bool) Nameserver {
	if !valid {
		return Nameserver{Addr: localhost(-1)}
	}
	return Nameserver{Addr: localhost(8899)}
}

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

// TODO: Saving this functional test
//func (s *NameserverSuite) TestResolve(c *C) {
//	ns = localServer(true)
//	ns.Connect()
//	go ns.Serve()
//
//	serv := []string{"127.0.0.1"}
//	r := dnsr.New(serv)
//	r.Servers[0] = "127.0.0.1:8899"
//
//	ips, err := r.LookupHost("www.thoughtworks.com")
//
//	c.Assert(err, IsNil)
//	c.Assert(ips, HasLen, 5)
//}

func (s *NameserverSuite) TestResponseIsReceived(c *C) {
	errChan := make(chan error)
	numChan := make(chan int)
	addrChan := make(chan *net.UDPAddr)

	ns = localServer(true)
	ns.Connect()

	conn, _ := net.ListenUDP("udp", localhost(8845))

	go func() {
		b := make([]byte, 512)
		n, ra, err := conn.ReadFromUDP(b)
		defer conn.Close()

		addrChan <- ra
		errChan <- err
		numChan <- n
	}()

	ns.handle(nil, localhost(8845))

	retAddr := <-addrChan
	c.Assert(retAddr.IP.String(), Equals, "127.0.0.1")
	c.Assert(retAddr.Port, Equals, 8899)
	c.Assert(<-errChan, IsNil)
	c.Assert(<-numChan, Equals, 5)
}
