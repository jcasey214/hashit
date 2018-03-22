package handler

import (
	"net/http"
	"encoding/json"
	"github.com/jcasey214/hashit/stats"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response, err := json.Marshal(stats.CurrentStats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
