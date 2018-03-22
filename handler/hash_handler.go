package handler

import (
	"net/http"
	"time"
	"strconv"
	"strings"
	"sync"
	"log"
	"github.com/jcasey214/hashit/hash"
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
			http.Error(w, "Missing required field in request", http.StatusBadRequest)
		} else {
			h.mutex.Lock()
			h.hashes = append(h.hashes, "")
			newIndex := len(h.hashes) - 1
			h.mutex.Unlock()
			go hashAndSave(pw, newIndex)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(strconv.Itoa(newIndex)))
		}
	} else {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}
}

func GetHash(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := strings.TrimPrefix(r.URL.Path, "/hash/")
		index, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		} else if index >= len(h.hashes) || index < 0 {
			http.Error(w, "Not Found", 404)
		} else {
			w.Write([]byte(h.hashes[index]))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func hashAndSave(password string, index int) {
	time.Sleep(5 * time.Second)
	log.Print("saving hash")
	h.hashes[index] = hash.Create(password)
}
