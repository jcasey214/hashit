package hash

import (
	"sync"
	"time"
	"log"
)

type Store struct {
	hashes []string
	mutex  sync.Mutex
}

var store = Store{hashes: []string{}}

func Save(i int, password string) {
	time.Sleep(5 * time.Second)
	log.Print("saving hash")
	store.hashes[i] = Create(password)
}

func GetId() int {
	store.mutex.Lock()
	store.hashes = append(store.hashes, "")
	newIndex := len(store.hashes) - 1
	store.mutex.Unlock()
	return newIndex
}

func Count() int {
	return len(store.hashes)
}

func Get(i int) string {
	return store.hashes[i]
}