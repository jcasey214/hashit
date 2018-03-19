package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/jcasey214/hashit/handler"
	"log"
	"time"
	"os/signal"
	"context"
)

func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	mux := http.NewServeMux()
	srv := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: mux}
	log.Printf("listening on port %s \n", port)

	mux.Handle("/hash", http.HandlerFunc(handler.HashHandler))
	mux.Handle("/shutdown", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		srv.Shutdown(ctx)
	}))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

	log.Println("Server gracefully stopped")
}
