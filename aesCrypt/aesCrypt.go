package aesCrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/pkg/errors"
	"io"
)

func encrypt(data, key []byte) ([]byte, error) {
	block, err := cifer(key)
	if err != nil {
		return nil, err
	}
	b := encodeStringFromBytes(data)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func EncryptToBase64(data string, key []byte) (string, error) {
	dataByte := []byte(data)
	res, err := encrypt(dataByte, key)
	if err != nil {
		return "", err
	}
	return encodeStringFromBytes(res), nil
}

func decrypt(data, key []byte) ([]byte, error) {
	block, err := cifer(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)
	return decodeBytes(string(data))
}

func DecryptFromBase64(data string, key []byte) (string, error) {
	dataByte, e := decodeBytes(data)
	if e != nil {
		return "", e
	}
	res, err := decrypt(dataByte, key)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func cifer(key []byte) (cipher.Block, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return c, nil
}
