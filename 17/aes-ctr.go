package _7

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func AesCtrEncrypt(key ,plainText []byte)([]byte,error){
	block,err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}

	vi := bytes.Repeat([]byte("1"),block.BlockSize())

	s := cipher.NewCTR(block,vi)

	dest := make([]byte,len(plainText))

	s.XORKeyStream(dest,plainText)

	return dest ,nil
}

func AesCtrDecrypt(key ,plainText []byte)([]byte,error){
	return AesCtrEncrypt(key,plainText)
}
