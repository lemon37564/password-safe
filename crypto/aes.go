package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AesEncrypt with 256 bits
func AesEncrypt(origData []byte, key []byte) ([]byte, error) {
	return aesEncrypt(origData, BytesToSha256(key), genIvFromPass(key), pkcs5Padding)
}

// AesDecrypt with 256 bits
func AesDecrypt(crypted []byte, key []byte) ([]byte, error) {
	return aesDecrypt(crypted, BytesToSha256(key), genIvFromPass(key), pkcs5UnPadding)
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

// Generate IV from password (16 bytes)
func genIvFromPass(pass []byte) (iv []byte) {
	tmp := BytesToSha256(pass)

	for i := 0; i < len(tmp); i++ {
		if i%4 == 0 {
			iv = append(iv, tmp[i])
		}
	}

	tmp = StrToSha256(string(tmp))

	for i := 0; i < len(tmp); i++ {
		if i%4 == 1 {
			iv = append(iv, tmp[i])
		}
	}

	// shuffle and change value
	primes := []int{2, 3, 5, 7, 11, 13, 17, 23, 29, 31, 37}
	j := int(iv[len(iv)-1]) % len(iv)
	for i := range iv {
		iv[i], iv[j] = iv[j], iv[i]
		j = (j + primes[i%len(primes)]) % len(iv)

		if i < j {
			tempInt := int(iv[i]) + primes[j%len(primes)]
			iv[i] = byte(tempInt % 256)
		}
	}

	return
}
