package crypto

import (
	"crypto/rand"
	"encoding/binary"
)

const (
	lower = "abcdefghijklmnopqrstuvwxyz"
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// more numbers
	number  = "01234567890123456789"
	special = "+-*/_~?!@#$%&><(){}[]"
	all     = lower + upper + number + special
)

// GeneratePassword with specific length
func GeneratePassword(length int) string {
	pass := ""
	for i := 0; i < length; i++ {
		pass += choose(all)
	}

	return pass
}

func choose(candidate string) string {
	number := make([]byte, 8)
	rand.Read(number)
	num := binary.BigEndian.Uint64(number)

	l := uint64(len(candidate))
	return string(candidate[num%l])
}
