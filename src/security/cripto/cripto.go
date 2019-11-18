package cripto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

func EncodeSHA256(secret []byte, data []byte) string{
	h := hmac.New(sha256.New, secret)
	h.Sum(data)
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func EncodeAES(key []byte, message string) (encmess string, err error) {
	plainText := []byte(message)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}

func DecodeAES(key []byte, securemess string) (decodedmess string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	decodedmess = string(cipherText)
	return
}