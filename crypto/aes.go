package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"log"
)

const ivLen = 16

// AesEncrypt with 256 bits
func AesEncrypt(origData []byte, key []byte, iv []byte) ([]byte, error) {
	return aesEncrypt(origData, BytesToSha256(key), iv, pkcs5Padding)
}

// AesDecrypt with 256 bits
func AesDecrypt(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	return aesDecrypt(crypted, BytesToSha256(key), iv, pkcs5UnPadding)
}

func aesEncrypt(origData []byte, key []byte, iv []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = paddingFunc(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func aesDecrypt(crypted, key []byte, iv []byte, unPaddingFunc func([]byte) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = unPaddingFunc(origData)
	return origData, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte("unpadding error")
	}
	return origData[:(length - unpadding)]
}

func GenerateIV() []byte {
	iv := make([]byte, ivLen)
	_, err := rand.Read(iv)
	if err != nil {
		log.Println(err)
		return nil
	}
	return iv
}
