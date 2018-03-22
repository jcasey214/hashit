package handler

import (
	"net/http"
	"strconv"
	"strings"
	"github.com/jcasey214/hashit/hash"
)

func CreateHash(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		pw := r.FormValue("password")

		if pw == "" {
			http.Error(w, "Missing required field in request", http.StatusBadRequest)
		} else {
			id := hash.GetId()
			go hash.Save(id, pw)

			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(strconv.Itoa(id)))
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
		} else if index >= hash.Count() || index < 0 {
			http.Error(w, "Not Found", 404)
		} else {
			w.Write([]byte(hash.Get(index)))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
