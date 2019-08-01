package main

import (
	"../../aesCrypt"
	//"bufio"
	//"fmt"
	//"github.com/mattermost/mattermost-bot-sample-golang/aesCrypt"
	//"log"
	"bufio"
	"fmt"
	"log"
	"os"
)
func main() {
	aesKey := os.Getenv("AES_KEY")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter password: ")
		password, _ := reader.ReadString('\n')
		password = password[:len(password)-2]
		aesPassword, e := aesCrypt.EncryptToBase64(password, []byte(aesKey))
		if e != nil {
			log.Println("Error encrypting password. " + e.Error())
		}
		println("Encrypted password: " + aesPassword)
	}
}
