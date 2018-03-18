package hash

import (
	"fmt"
	"crypto/sha512"
	"encoding/base64"
)

func Hash(password string) string {
	fmt.Println("hashing password")
	pw := []byte(password)
	hasher := sha512.New()
	hasher.Write(pw)
	digest := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return digest
}
