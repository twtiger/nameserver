package main

import (
	"fmt"
	"log"
)

func initLogger() {
	flags := log.Ldate | log.Ltime | log.Llongfile
	log.SetFlags(flags)
	log.SetPrefix("[logger:] ")
}

func main() {
	initLogger()

	err := run()
	if err != nil {
		errMsg := fmt.Sprintf("Error in starting dns nameserver")
		log.Printf(errMsg)
	}
}
