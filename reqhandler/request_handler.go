package reqhandler

import (
	"fmt"
	"net"
)

func HandleConnection(net.Conn) bool {
	fmt.Printf("in connection")
	return true
}
