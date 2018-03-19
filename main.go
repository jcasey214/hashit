package main

import (
	"net/http"
	"os"
	"github.com/jcasey214/hashit/handler"
	"log"
	"context"
	"time"
	"fmt"
)

func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	done := make(chan bool)

	srv := &http.Server{Addr: fmt.Sprintf(":%s", port)}
	log.Printf("listening on port %s \n", port)

	http.HandleFunc("/hash", http.HandlerFunc(handler.HashHandler))
	http.HandleFunc("/shutdown", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		srv.Shutdown(ctx)
	}))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-done
}
