package nameserver

import (
	"fmt"
	"log"
	"net"
)

func initLogger() {
	flags := log.Ldate | log.Ltime | log.Llongfile
	log.SetFlags(flags)
	log.SetPrefix("[logger:] ")
}

func main() {
	initLogger()
	n := Nameserver{Addr: &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8853,
	}}
	err := n.Connect()
	if err != nil {
		errMsg := fmt.Sprintf("Error in starting dns nameserver: %s", err.Error())
		log.Printf(errMsg)
	}
}
