package handler

import (
	"net/http"
	"fmt"
	"time"
	"github.com/jcasey214/hashit/hash"
	"strconv"
)

var hashes []string

func HashHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if r.Method == "POST" {
		r.ParseForm()
		pw := r.FormValue("password")

		if pw == "" {
			http.Error(w, "Missing required field in request", 400)
		} else {
			hashes = append(hashes, "")
			newIndex := len(hashes) - 1
			w.Header().Set("Content-Type", "text/plain")
			go hashPassword(pw, newIndex)
			w.Write([]byte(strconv.Itoa(newIndex)))
		}
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func hashPassword(password string, index int) {
	time.Sleep(5*time.Second)
	hashes[index] = hash.Hash(password)
	fmt.Println(hashes)
}
