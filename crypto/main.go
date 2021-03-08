// +build ignore

package main

import (
	"crypto/rand"
	"fmt"
	"pass-safe/crypto"
	"time"
)

func match(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {

	key := []byte("my-key")
	message := make([]byte, 1048576)
	iv := crypto.GenerateIV()
	rand.Read(message)

	t1 := time.Now()
	encrypted, err := crypto.AesEncrypt(message, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println("encrypt spent:", time.Since(t1))

	time.Sleep(time.Second * 15)

	t1 = time.Now()
	decrypted, err := crypto.AesDecrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println("decrypt spent:", time.Since(t1))

	if match(message, decrypted) {
		fmt.Println("matched")
	} else {
		fmt.Println("not matched")
	}
}
