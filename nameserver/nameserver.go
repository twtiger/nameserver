package nameserver

import (
	"fmt"
	"net"
)

// Nameserver handles DNS queries
type Nameserver struct {
	Addr *net.UDPAddr
	ucon *net.UDPConn
}

// Connect will begin listening for DNS queries on localhost:8853
func (n *Nameserver) Connect() error {
	c, err := net.ListenUDP("udp", n.Addr)
	if err != nil {
		return fmt.Errorf("unable to connect: %s", err.Error())
	}
	n.ucon = c
	return nil
}

// Serve will begin listening for DNS queries and responses
func (n *Nameserver) Serve() error {
	if n.ucon == nil {
		return fmt.Errorf("not connected: must successfully connect with nameserver.Connect first")
	}
	defer n.teardown()
	for {
		bs := make([]byte, 512)
		_, ra, err := n.ucon.ReadFromUDP(bs)
		if err != nil {
			return err
		}
		n.handle(bs, ra)
	}
}

func (n *Nameserver) teardown() error {
	return n.ucon.Close()
}

func (n *Nameserver) handle(b []byte, ra *net.UDPAddr) {
	// TODO: Waiting on interface agreement
	//msg, err := m.Unpack(b)
	//if err != nil {
	//	//Unable to unmarshal
	//}
	//res, err := r.Respond(msg)
	//if err != nil {
	//	return err
	//}
	//// Do compression
	//p, err := m.Pack(res)
	//if err != nil {
	//	return err
	//}

	p := []byte("hello")
	_, err := n.ucon.WriteTo(p, ra)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
}
