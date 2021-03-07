package manager

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pass-safe/crypto"
	"strings"
)

const docName = ".enc"

type Manager struct {
	data    map[string][]byte
	keyHash []byte
	file    *os.File
}

func NewManager() *Manager {
	m := new(Manager)

	info, err := os.Stat(docName)
	if err != nil {
		m.createDoc(notFound)
	}

	m.readDoc(info)

	return m
}

func (m Manager) Loop() {

	defer m.file.Close()
	defer m.dump()

	fmt.Println()
	fmt.Println("Password Manager V0.0.1")
	log.Println()
	fmt.Println("Type \"help\" for more information")
	fmt.Println()

	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		cmd, err := input.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// eliminate '\n'
		cmd = cmd[:len(cmd)-2]
		if cmd == "" {
			continue
		}

		cmds := strings.Split(cmd, " ")
		maincmd := strings.ToLower(cmds[0])

		switch maincmd {
		case "help":
			printHelp()
		case "list":
			printList()
		case "show":
			passof(cmds[1:])
		case "set":
			m.setpass(cmds[1:])
		case "delete":
			deldata(cmds[1:])
		case "genpass":
			genpass(cmds[1:])
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Unkown command. Try \"help\"")
		}

		fmt.Println()
	}
}

// print all available commands
func printHelp() {
	fmt.Println("This is a password management tool.")
	fmt.Println("The encryption algorithm is based on AES-256")
	fmt.Println("Note: it is highly suggested to back up your", docName, "file")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Println("help                list all the commands")
	fmt.Println("list                list all the names")
	fmt.Println("show NAME           show the password of NAME")
	fmt.Println("set NAME PASS       set PASS as password of NAME")
	fmt.Println("delete NAME         delete the data NAME")
	fmt.Println("exit                terminate the program")
}

// print all the keys(not password)
func printList() {

}

// get the password of the key
// -c copy the password into clipboard
// -h don't show the password(need to use with -c)
func passof(cmds []string) {
	var key string
	fmt.Print("Enter user's password to confirm: ")
	fmt.Scanln(&key)

}

// set the password of the key
func (m *Manager) setpass(cmds []string) {
	if len(cmds) == 2 {
		_, exi := m.data[cmds[0]]
		if exi {
			var c rune
			fmt.Print(cmds[0], "has already exists, do you want to overwrite? [y/n] ")
			fmt.Scanln(&c)
			if c == 'y' {
				// m.data[cmds[0]] = crypto.AesEncrypt([]byte())
			}
		}
	}
}

// delete one password or key
// -all delete all (need confirm)
func deldata(cmds []string) {
	if len(cmds) == 0 {
		fmt.Println("missing arguments.")
	} else if len(cmds) > 1 {
		fmt.Println("unknown argument:", cmds[1])
	} else {
		if cmds[0] == "-all" {

		} else {

		}
	}
}

// generate one reliable password with length 16
// -s=NAME set the name as the key generated
// -h don't show the password
func genpass(cmds []string) {
	set := false
	hide := false

	for _, v := range cmds {
		if v[0] == '-' {
			if v[1] == 'h' {
				hide = true
			} else if v[1] == 's' {
				set = true
				for i := 2; i < len(v); i++ {

				}
			} else {
				fmt.Println("unknown argument: -", string(v[1]))
			}
		}
	}

	password := crypto.GenPass()
	if !hide {
		fmt.Println(password)
	}
	if set {
		fmt.Println("set")
	}
}
