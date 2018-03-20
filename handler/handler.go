package handler

import (
	"net/http"
	"time"
	"github.com/jcasey214/hashit/hash"
	"strconv"
	"strings"
)

var hashes []string

func CreateHash(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		pw := r.FormValue("password")

		if pw == "" {
			http.Error(w, "Missing required field in request", 400)
		} else {
			hashes = append(hashes, "")
			newIndex := len(hashes) - 1
			w.Header().Set("Content-Type", "text/plain")
			go hashAndSave(pw, newIndex)
			w.Write([]byte(strconv.Itoa(newIndex)))
		}
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func GetHashById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := strings.TrimPrefix(r.URL.Path, "/hash/")
		index, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, "Bad Request", 400)
		} else if index >= len(hashes) {
			http.Error(w, "Not Found", 404)
		} else {
			w.Write([]byte(hashes[index]))
		}

	} else {
		http.Error(w, "Invalid request method", 405)
	}
}

func hashAndSave(password string, index int) {
	time.Sleep(5 * time.Second)
	hashes[index] = hash.Hash(password)
}
