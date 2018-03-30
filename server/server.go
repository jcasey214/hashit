package server

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"context"
	"github.com/jcasey214/hashit/handler"
	"github.com/jcasey214/hashit/stats"
)

func Run(port int) chan bool {
	srv := http.Server{Addr: fmt.Sprintf(":%d", port)}
	log.Printf("listening on port %d \n", port)

	doneChan := make(chan bool)

	http.HandleFunc("/hash", stats.Recorder(handler.CreateHash))
	http.HandleFunc("/hash/", http.HandlerFunc(handler.GetHash))
	http.HandleFunc("/stats", http.HandlerFunc(handler.GetStats))
	http.HandleFunc("/shutdown", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		srv.Shutdown(ctx)
		w.WriteHeader(http.StatusOK)
		doneChan <- true
	}))

	go func(){
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	return doneChan
}
