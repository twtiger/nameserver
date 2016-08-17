package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8853")
	defer conn.Close()
	if err != nil {
		fmt.Println("cannot make connection")
	}
	for i := 0; i < 100; i++ {
		fmt.Println("CONNECTED")
		msg := []byte("Hello world!\n")
		b, err := conn.Write(msg)
		if err != nil {
			fmt.Printf("cannot write to connection :( %s\n", err)
		} else {
			fmt.Printf("number of bytes written: %d\n", b)
		}
	}
}
