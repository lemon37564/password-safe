package manager

import (
	"fmt"
	"os"
	"pass-safe/crypto"
)

const (
	notFound = iota
	broken
)

func (m *Manager) createDoc(stat int) {
	if stat == notFound {
		fmt.Println(docName + " file not found, initializing...")
	} else {
		fmt.Println(docName + " has broken, initailzing...")
	}

	f, err := os.Create(docName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Done. Please type your user password: ")
	var pw string
	fmt.Scanln(&pw)

	m.keyHash = crypto.StrToSha256(pw)

	fmt.Println("Success.")
}

func (m *Manager) readDoc(info os.FileInfo) {
	f, err := os.Open(docName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// check the file is noraml
	if info.Size() < 32 {
		m.createDoc(broken)
	}

	data := make([]byte, info.Size())
	f.Read(data)

	m.parse(data)
}
