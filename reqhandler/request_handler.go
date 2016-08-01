package reqhandler

import (
	"fmt"
	"net"
)

func handleConnection(net.Conn) bool {
	fmt.Printf("in connection")
	return true
}
