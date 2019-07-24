package main

import (
	"bufio"
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/aesCrypt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter key (32 characters): ")
	aesKey, _ := reader.ReadString('\n')
	aesKey = aesKey[:len(aesKey) - 1]

	for {
		fmt.Print("Enter database password: ")
		password, _ := reader.ReadString('\n')
		aesPass := aesCrypt.EncryptToBase64(password, []byte(aesKey))
		fmt.Println("AES password: " + aesPass)
	}
}