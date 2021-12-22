package cripto_test

import (
	"fmt"
	"testing"

	"github.com/guru-invest/guru.framework/src/security/cripto"
)

func TestSHA256Creation(t *testing.T) {
	secret := "secret"
	data := "teste de informação criptografada"
	fmt.Println(cripto.EncodeSHA256([]byte(secret), []byte(data)))
}

func TestAESCreation(t *testing.T) {
	secret := "ProjetoGuru@@abc"
	data := "teste de informação criptografada"
	hash, err := cripto.EncodeAES([]byte(secret), data)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(hash)
}

func TestAESDecode(t *testing.T) {
	secret := "ProjetoGuru@@abc"
	hash := "wc9GNxcKcdg5LjLlTEpIiA9ve5L-zrOS6zh6IgC9D1GWT52SjlKV6uhZrvRzLr04Rw4S"
	phrase, err := cripto.DecodeAES([]byte(secret), hash)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(phrase)
}

func TestAESDecodeShortCiphertext(t *testing.T) {
	secret := "Projeto@@abGuruc"
	hash := ""
	_, err := cripto.DecodeAES([]byte(secret), hash)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\033[1;31m%s\033[0m", "Ciphertext block size must be too short!\n")
	t.Fail()
}
