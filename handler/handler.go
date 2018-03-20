package handler

import (
	"net/http"
	"time"
	"github.com/jcasey214/hashit/hash"
	"strconv"
	"strings"
	"sync"
)

type HashStore struct {
	hashes []string
	mutex  sync.Mutex
}

var h = HashStore{hashes: []string{}}

func CreateHash(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		pw := r.FormValue("password")

		if pw == "" {
			http.Error(w, "Missing required field in request", 400)
		} else {
			h.mutex.Lock()
			h.hashes = append(h.hashes, "")
			newIndex := len(h.hashes) - 1
			h.mutex.Unlock()
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
		} else if index >= len(h.hashes) {
			http.Error(w, "Not Found", 404)
		} else {
			w.Write([]byte(h.hashes[index]))
		}

	} else {
		http.Error(w, "Invalid request method", 405)
	}
}

func hashAndSave(password string, index int) {
	time.Sleep(5 * time.Second)
	h.hashes[index] = hash.Hash(password)
}
