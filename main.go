package dnsserv

import "fmt"

func main() {
	err := start()
	if err != nil {
		fmt.Println("getting error")
	}
}
