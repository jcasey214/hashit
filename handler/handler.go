package handler

import (
	"net/http"
	"fmt"
	"time"
	"github.com/jcasey214/hashit/hash"
)

func HashHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if r.Method == "POST" {
		pw := r.FormValue("password")

		if pw == "" {
			http.Error(w, "Missing required field in request", 400)
		} else {
			time.Sleep(5 * time.Second)
			w.Header().Set("Content-Type", "text/plain")
			r.ParseForm()
			w.Write([]byte(hash.Hash(pw)))
		}
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}
