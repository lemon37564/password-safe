// +build ignore

package main

import (
	"fmt"
	"pass-safe/crypto"
)

func main() {
	key := []byte("my-key")
	orig := []byte("this is secret message.")
	fmt.Println(orig)

	encrypted, err := crypto.AesEncrypt(orig, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(encrypted)

	decrypted, err := crypto.AesDecrypt(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decrypted))
}
