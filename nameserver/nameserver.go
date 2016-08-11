package nameserver

import (
	"log"
	"net"
)

const dnsMsgSize = 512
const port = 8853

func getUDPAddr() *net.UDPAddr {
	return &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: port,
	}
}

// Start will begin listening for DNS queries on port 8853
func Start() error {
	udpConn, err := net.ListenUDP("udp", getUDPAddr())
	if err != nil {
		return err
	}

	for {
		b := make([]byte, dnsMsgSize)
		_, _, err := udpConn.ReadFrom(b)
		if err != nil {
			log.Printf("Error in reading message " + err.Error())
		}
		return err
	}
}
