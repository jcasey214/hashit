package stats

import (
	"sync"
	"net/http"
	"log"
	"time"
	"strconv"
)

type ServerStats struct {
	Total         int `json:"total"`
	Average       int `json:"average"`
	totalDuration int
	mutex         sync.Mutex
}

var CurrentStats = ServerStats{Total: 0, Average: 0}

func updateStats(duration int) {
	CurrentStats.mutex.Lock()
	CurrentStats.Total += 1
	CurrentStats.totalDuration += duration
	CurrentStats.Average = CurrentStats.totalDuration / CurrentStats.Total
	CurrentStats.mutex.Unlock()
}

func Stats(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := makeTimestamp()
		log.Println("Before")
		h.ServeHTTP(w, r) // call original
		log.Println("After")
		stop := makeTimestamp()
		defer log.Print(strconv.Itoa(int(stop - start)))
	})
}

func makeTimestamp() int64 {
	return time.Now().UnixNano()
}
