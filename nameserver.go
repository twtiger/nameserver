package nameserver

import (
	"errors"
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
		return errors.New("not connected: must successfully connect with nameserver.Connect first")
	}
	defer n.teardown()
	for {
		b := make([]byte, 512)
		_, retAddr, err := n.ucon.ReadFromUDP(b)
		if err != nil {
			return err
		}
		serializedResponse := respondTo(b)
		n.reply(serializedResponse, retAddr)
	}
}

func (n *Nameserver) reply(serializedResponse []byte, retAddr *net.UDPAddr) {
	_, err := n.ucon.WriteTo(serializedResponse, retAddr)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
}

func (n *Nameserver) teardown() error {
	return n.ucon.Close()
}

func respondTo(b []byte) []byte {
	msg := &message{}
	_ = msg.deserialize(b)
	sr, _ := msg.serialize()
	return sr
}
