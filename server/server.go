package server

import (
	"net/http"
	"fmt"
	"log"
	"github.com/jcasey214/hashit/handler"
	"time"
	"context"
)

func CreateServer(port string) chan bool {
	srv := &http.Server{Addr: fmt.Sprintf(":%s", port)}
	log.Printf("listening on port %s \n", port)

	done := make(chan bool)
	http.HandleFunc("/hash", http.HandlerFunc(handler.HashHandler))
	http.HandleFunc("/shutdown", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		srv.Shutdown(ctx)
		done <- true
	}))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	return done
}