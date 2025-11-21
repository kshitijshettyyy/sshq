package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

var key = []byte("this_is_a_demo_key_32bytes_long!") // replace later

func Encrypt(data []byte) []byte {
	block, _ := aes.NewCipher(key)
	aesGCM, _ := cipher.NewGCM(block)
	nonce := make([]byte, aesGCM.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := aesGCM.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(ciphertext []byte) []byte {
	block, _ := aes.NewCipher(key)
	aesGCM, _ := cipher.NewGCM(block)
	nonceSize := aesGCM.NonceSize()
	nonce, text := ciphertext[:nonceSize], ciphertext[nonceSize:]
	data, _ := aesGCM.Open(nil, nonce, text, nil)
	return data
}
