package main

import (
	"flag"
	"github.com/jcasey214/hashit/server"
)

func main() {
	port := flag.Int("port", 8080, "port to launch server on")

	flag.Parse()

	doneChan := server.Run(*port)

	<-doneChan
}
