package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/jcasey214/hashit/handler"
	"log"
)

func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}
	fmt.Printf("listening on port %s \n", port)

	http.HandleFunc("/hash", handler.HashHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
