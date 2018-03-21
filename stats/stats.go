package stats

import (
	"sync"
	"net/http"
	"time"
)

type ServerStats struct {
	Total         int     `json:"total"`
	Average       float32 `json:"average"`
	totalDuration float32
	mutex         sync.Mutex
}

var CurrentStats = ServerStats{Total: 0, Average: 0}

func Recorder(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := makeTimestamp()

		h.ServeHTTP(w, r)

		stop := makeTimestamp()
		elapsed := float32(stop-start) / float32(time.Millisecond)

		updateStats(elapsed)
	})
}

func updateStats(duration float32) {
	CurrentStats.mutex.Lock()
	CurrentStats.Total += 1
	CurrentStats.totalDuration += duration
	CurrentStats.Average = CurrentStats.totalDuration / float32(CurrentStats.Total)
	CurrentStats.mutex.Unlock()
}

func makeTimestamp() int64 {
	return time.Now().UnixNano()
}
