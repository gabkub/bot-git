package aesCrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/mattermost/mattermost-bot-sample-golang/logs"
	"log"
)

func encrypt(data, key []byte) []byte {
	return nil
}

func EncryptToBase64(data string, key []byte) string {
	return ""
}

func decrypt(data, key []byte) []byte {
	return nil
}

func DecryptFromBase64(data string, key []byte) string {
	return ""
}

func cifer(key []byte) cipher.Block {
	c, err := aes.NewCipher(key)
	if err != nil {
		logs.WriteToFile("Error ciphering the aes key.")
		log.Fatal("Error ciphering the aes key.")
	}
	return c
}
