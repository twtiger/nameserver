package nameserver

import (
	"net"
	"testing"

	dnsr "github.com/bogdanovich/dns_resolver"
	. "gopkg.in/check.v1"
)

const nsPort = 8899

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
	return Nameserver{Addr: localhost(nsPort)}
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

//TODO: Saving this functional test
// func (s *NameserverSuite) TestResolve(c *C) {
// 	ns = localServer(true)
// 	ns.Connect()
// 	go ns.Serve()

// 	serv := []string{"127.0.0.1"}
// 	r := dnsr.New(serv)
// 	r.Servers[0] = "127.0.0.1:8899"

// 	ips, err := r.LookupHost("twtiger.com")

// 	c.Assert(err, IsNil)
// 	c.Assert(ips, HasLen, 2)
// }

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

func createRecordNameInBytesForTwtiger() []byte {
	recordName := []byte{7}
	recordName = append(recordName, []byte("twtiger")...)
	recordName = append(recordName, 3)
	recordName = append(recordName, []byte("com")...)
	recordName = append(recordName, 0)
	return recordName
}

func (s *NameserverSuite) Test_CreationOfSerializedResponseFromQuery(c *C) {
	header := make([]byte, 12)
	recordName := createRecordNameInBytesForTwtiger()

	message := header
	message = append(message, recordName...)

	response := respondTo(message)

	c.Assert(response[0:12], DeepEquals, header)
	c.Assert(response[12:25], DeepEquals, recordName)

	// TODO waiting for serialize to be completed

	// recordType := []byte{0,1}
	// recordClass := []byte{0,1}
	// recordTTL := []byte{0,0,14,16}
	// recordRDLength := []byte{4}
	// recordRData := []byte{123, 123, 7, 8}

	// c.Assert(response[25:27], DeepEquals, recordType) // qtype
	// c.Assert(response[27:29], DeepEquals, recordClass) // qclass
	// c.Assert(response[29:42], DeepEquals, recordName)
	// c.Assert(response[42:44], DeepEquals, recordType)
	// c.Assert(response[44:46], DeepEquals, recordClass)
	// c.Assert(response[46:48], DeepEquals, recordTTL)
	// c.Assert(response[48], DeepEquals, recordRDLength)
	// c.Assert(response[49:50], DeepEquals, recordRData)
}
