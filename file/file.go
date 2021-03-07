package file

import (
	"encoding/json"
	"log"
	"os"
)

const docName = "safe"

// Create file to store encrypted data
func Create() {
	f, err := os.Create(docName)
	if err != nil {
		log.Println("creating file:", err)
		return
	}
	defer f.Close()
}

// Store the encrypted data into file with json
func Store(data map[string][]string, key []byte) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Println("marshaling object:", err)
		return
	}

	// encrypt here
	enc := j

	store(enc)
}

func store(data []byte) {
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

// Read the encrypted data from file
func Read() (m map[string][]string) {

	data := read()

	// decrypt here
	dec := data

	err := json.Unmarshal(dec, &m)
	if err != nil {
		log.Println("unmarshaling object:", err)
		return
	}

	return
}

func read() (data []byte) {
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
