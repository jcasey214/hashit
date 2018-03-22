package hash

import (
	"crypto/sha512"
	"encoding/base64"
)

func Create(password string) string {
	pw := []byte(password)
	hasher := sha512.New()
	hasher.Write(pw)
	digest := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return digest
}
