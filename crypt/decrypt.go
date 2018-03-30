package crypt

import (
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Decrypt(cryptoText []byte, secretKey string) string {
	key := []byte(secretKey)
	ciphertext, _ := base64.URLEncoding.DecodeString(string(cryptoText))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
