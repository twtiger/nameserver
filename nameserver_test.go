package nameserver

import (
	"encoding/binary"
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
// func (s *NameserverSuite) Test_ReceivesValidResponseForAuthZoneAddress(c *C) {
// const authZoneAddr := "twtiger.com"

// 	ns = localServer(true)
// 	ns.Connect()
// 	go ns.Serve()

// 	serv := []string{"127.0.0.1"}
// 	r := dnsr.New(serv)
// 	r.Servers[0] = "127.0.0.1:8899"

// 	ips, err := r.LookupHost(authZoneAddr)

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

func recordNameForSecondLevelDomain(firstLevelDom string, secondLevelDom string) []byte {
	lenfirstLevelDom := byte(len(firstLevelDom))
	lenSecondLevelDom := byte(len(secondLevelDom))

	recordName := []byte{lenfirstLevelDom}
	recordName = append(recordName, []byte(firstLevelDom)...)
	recordName = append(recordName, lenSecondLevelDom)
	recordName = append(recordName, []byte(secondLevelDom)...)
	recordName = append(recordName, 0)
	return recordName
}

func qtypeAndQclass() (qtype []byte, qclass []byte) {
	b := make([]byte, 2) // TODO pull into a test helper file
	binary.BigEndian.PutUint16(b, uint16(qtypeA))
	qtype = append(qtype, b...)
	b = make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(qclassIN))
	qclass = append(qclass, b...)
	return
}

func (s *NameserverSuite) Test_CreationOfSerializedResponseFromQuery(c *C) {
	header := make([]byte, 12)
	recordName := recordNameForSecondLevelDomain("twtiger", "com")
	message := append(header, recordName...)
	qtype, qclass := qtypeAndQclass()
	message = append(message, qtype...)
	message = append(message, qclass...)

	response := respondTo(message)

	c.Assert(response[0:12], DeepEquals, header)
	c.Assert(response[12:25], DeepEquals, recordName) //qname

	//recordType := []byte{0, 1}
	//recordClass := []byte{0, 1}
	//recordTTL := []byte{0, 0, 14, 16}
	//recordRDLength := []byte{4}
	//recordRData := []byte{123, 123, 7, 8}

	//c.Assert(response[25:27], DeepEquals, qtype)
	//c.Assert(response[27:29], DeepEquals, qclass)
	//c.Assert(response[29:42], DeepEquals, recordName) // answers begin
	//c.Assert(response[42:44], DeepEquals, recordType)
	//c.Assert(response[44:46], DeepEquals, recordClass)
	//c.Assert(response[46:48], DeepEquals, recordTTL)
	//c.Assert(response[48], DeepEquals, recordRDLength)
	//c.Assert(response[49:50], DeepEquals, recordRData)
}

// TODO waiting for serialize to be completed
// var resBchan = make(chan []byte)
// var resReturnAddrChan = make(chan *net.UDPAddr)
// var resErrChan = make(chan error)

// func setupResolver(setupAddr *net.UDPAddr, bmsg []byte, sendTo *net.UDPAddr) {
// 	listener, _ := net.ListenUDP("udp", setupAddr)
// 	go func() {
// 		listener.WriteTo(bmsg, sendTo)
// 		b := make([]byte, 512)
// 		_, ra, err := listener.ReadFromUDP(b)
// 		defer listener.Close()

// 		resBchan <- b
// 		resReturnAddrChan <- ra
// 		resErrChan <- err
// 	}()
// }

// func (s *NameserverSuite) Test_ReceivesValidResponseForExtZoneAddress(c *C) {
// 	recordName := recordNameForSecondLevelDomain("wireshark", "org")
// 	message := append(make([]byte, 12), recordName...)

// 	ns = localServer(true)
// 	ns.Connect()
// 	go ns.Serve()

// 	const resolverPort = 8866
// 	setupResolver(localhost(resolverPort), message, localhost(nsPort))

// 	respToRes := <-resBchan
// 	respRetAddr := <-resReturnAddrChan
// 	respErr := <-resErrChan
// 	c.Assert(respToRes, DeepEquals, message)
// 	c.Assert(respRetAddr.IP.String(), Equals, "127.0.0.1")
// 	c.Assert(respRetAddr.Port, Equals, nsPort)
// 	c.Assert(respErr, IsNil)
// }
