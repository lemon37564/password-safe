package storage

import (
	"encoding/json"
	"log"
	"os"
	"pass-safe/crypto"
)

const docName = "pass.safe"

// Create file to store encrypted data
func Create() {
	f, err := os.Create(docName)
	if err != nil {
		log.Println("creating file:", err)
		return
	}
	defer f.Close()
}

func store(data map[string]Pair, key []byte) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Println("marshaling object:", err)
		return
	}

	// encrypt here
	enc, err := crypto.AesEncrypt(j, key)
	if err != nil {
		log.Println("aes encrypting:", err)
	}

	storeToFile(enc)
}

func storeToFile(data []byte) {
	f, err := os.Create(docName)
	if err != nil {
		log.Println("opening document:", err)
		return
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		log.Println("writing data:", err)
		return
	}
}

func read(key []byte) (m map[string]Pair) {
	m = make(map[string]Pair)

	data := readFromFile()
	if len(data) == 0 {
		Create()
		return
	}

	// decrypt here
	dec, err := crypto.AesDecrypt(data, key)
	if err != nil {
		log.Println("aes decrypting:", err)
		return
	}

	err = json.Unmarshal(dec, &m)
	// if this error occurs then probably the key is wrong
	if err != nil {
		log.Println("unmarshaling object:", err)
		return
	}

	return
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
