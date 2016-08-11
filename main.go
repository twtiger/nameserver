package main

import (
	"fmt"
	"log"

	"github.com/twtiger/toy-dns-nameserver/nameserver"
)

func initLogger() {
	flags := log.Ldate | log.Ltime | log.Llongfile
	log.SetFlags(flags)
	log.SetPrefix("[logger:] ")
}

func main() {
	initLogger()

	err := nameserver.Start()
	if err != nil {
		errMsg := fmt.Sprintf("Error in starting dns nameserver: %s", err.Error())
		log.Printf(errMsg)
	}
}
