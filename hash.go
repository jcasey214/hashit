package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println(Hash(os.Args[1]))
	} else {
		fmt.Println("Please pass in a single password to hash")
	}
}

func Hash(password string) string {
	pw := []byte(password)
	hasher := sha512.New()
	hasher.Write(pw)
	digest := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return digest
}
