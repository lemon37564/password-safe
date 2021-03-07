package crypto

import "crypto/sha256"

func StrToSha256(key string) []byte {
	salt := "lemonteaIsTheBest"
	key += salt

	hasher := sha256.New()
	hasher.Write([]byte(key))

	return hasher.Sum(nil)
}

func BytesToSha256(key []byte) []byte {
	salt := []byte("lemonteaIsTheBest")
	key = append(key, salt...)

	hasher := sha256.New()
	hasher.Write([]byte(key))

	return hasher.Sum(nil)
}
