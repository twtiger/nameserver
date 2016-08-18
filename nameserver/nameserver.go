package nameserver

import (
	"fmt"
	"net"
)

const dnsMsgSize = 512
const port = 8853

// Nameserver handles DNS queries
type Nameserver struct {
	Addr *net.UDPAddr
	pcon *net.UDPConn
}

// Connect will begin listening for DNS queries on localhost:8853
func (n *Nameserver) Connect() error {
	c, err := net.ListenUDP("udp", n.Addr)
	if err != nil {
		return fmt.Errorf("unable to connect: %s", err.Error())
	}
	n.pcon = c
	return nil
}

// Serve will begin listening for DNS queries and responses
func (n *Nameserver) Serve() error {
	if n.pcon == nil {
		return fmt.Errorf("not connected: must successfully connect with nameserver.Connect first")
	}
	defer n.teardown()
	return nil
}

func (n *Nameserver) teardown() error {
	return n.pcon.Close()
}
