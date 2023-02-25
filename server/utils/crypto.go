package utils

import (
	"github.com/deatil/go-cryptobin/cryptobin/crypto"
)

func Crypto(str string) string {
	return crypto.
		FromString(str).
		SetKey("dfertf12dfertf12").
		Aes().
		ECB().
		PKCS7Padding().
		Encrypt().
		ToBase64String()
}

func Decryption(str string) string {
	return crypto.
		FromBase64String(str).
		SetKey("dfertf12dfertf12").
		Aes().
		ECB().
		PKCS7Padding().
		Decrypt().
		ToString()
}
