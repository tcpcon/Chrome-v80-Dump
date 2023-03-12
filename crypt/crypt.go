package crypt 

import (
	"crypto/cipher"
	"crypto/aes"
)

func Aes256GcmDecrypt(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	dec, err := gcm.Open(nil, data[3:15], data[15:], nil)
	if err != nil {
		panic(err)
	}

	return dec
}
