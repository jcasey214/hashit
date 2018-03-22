package main

import (
	"os"
	"github.com/jcasey214/hashit/server"
)

func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	doneChan := server.Run(port)

	<-doneChan
}
