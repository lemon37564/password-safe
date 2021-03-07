package crypto

import (
	"crypto/rand"
)

// len(chars) = 64
// l and I is removed because it's similar to 1
// O and o is removed because it's similar to 0
// q is removed because sometimes similar to 9
const chars = "abcdefghijkmnprstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ0123456789+-=@!%?"
const clen = len(chars)

// generate password of length 16
func GenPass() string {
	pass := make([]byte, 16)
	// cryptographically secured bytes
	rand.Read(pass)

	// 256 % clen == 0 so it is fair to every char
	for i := range pass {
		pass[i] = chars[int(pass[i])%clen]
	}

	return string(pass)
}
