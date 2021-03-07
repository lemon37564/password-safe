package manager

import (
	"fmt"
	"os"
)

func (m *Manager) parse(info []byte) {
	keyHash := info[:32]
	m.keyHash = keyHash
}

func (m Manager) dump() {
	fmt.Println("dumped")
	f, err := os.Create(docName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(m.keyHash)
}
