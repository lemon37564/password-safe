package storage

import (
	"encoding/json"
	"log"
	"os"
	"pass-safe/crypto"
)

/////////////////////////////////////////////////////////////
//// document format: 0~15 bytes-> iv, 16~ bytes -> body ////
/////////////////////////////////////////////////////////////

const (
	docName = "pass.safe"
	ivLen   = 16 // bytes
)

// KeyError represents aes key error
type KeyError struct {
	s string
}

func (k KeyError) Error() string {
	return k.s
}

// Create file to store encrypted data
func Create() {
	f, err := os.Create(docName)
	if err != nil {
		log.Println("creating file:", err)
		return
	}
	defer f.Close()
}

// IsFileExist return boolean value that tells file exist or not
func IsFileExist() bool {
	_, err := os.Stat(docName)
	return err == nil
}

func store(data map[string]Pair, key []byte, iv []byte) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Println("marshaling object:", err)
		return
	}

	// encrypt here
	enc, err := crypto.AesEncrypt(j, key, iv)
	if err != nil {
		log.Println("aes encrypting:", err)
	}

	storeToFile(iv, enc)
}

func storeToFile(iv []byte, data []byte) {
	f, err := os.Create(docName)
	if err != nil {
		log.Println("opening document:", err)
		return
	}
	defer f.Close()

	_, err = f.Write(iv)
	if err != nil {
		log.Println("writing data:", err)
		return
	}

	_, err = f.Write(data)
	if err != nil {
		log.Println("writing data:", err)
		return
	}
}

func read(key []byte) (map[string]Pair, error) {
	m := make(map[string]Pair)

	data := readFromFile()
	if len(data) == 0 || len(data) <= ivLen {
		return m, nil
	}

	iv := data[0:ivLen]
	data = data[ivLen:]
	// decrypt here
	dec, err := crypto.AesDecrypt(data, key, iv)
	if err != nil {
		log.Println("aes decrypting:", err)
		return m, nil
	}

	err = json.Unmarshal(dec, &m)
	// if this error occurs then probably the key is wrong
	if err != nil {
		log.Println("unmarshaling object:", err)
		return m, KeyError{s: "Wrong Password or damaged file"}
	}

	return m, nil
}

func readFromFile() (data []byte) {
	f, err := os.Open(docName)
	if err != nil {
		log.Println("opening document:", err)
		return
	}
	defer f.Close()

	info, err := os.Stat(docName)
	if err != nil {
		log.Println("stating document:", err)
		return
	}

	data = make([]byte, info.Size())
	_, err = f.Read(data)
	if err != nil {
		log.Println("reading data:", err)
		return
	}

	return
}
