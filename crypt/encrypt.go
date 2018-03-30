package crypt

import (
	"io"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"crypto/rand"
)

func Encrypt(file []byte, secretKey string) string {
	key := []byte(secretKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(file))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], file)

	return base64.URLEncoding.EncodeToString(ciphertext)
}
