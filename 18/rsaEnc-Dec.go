package _8

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

// 加密
func RsaPubKeyEncryptData(plainText []byte, filename string) ([]byte, error) {
	// 获取公钥进行加密
	pubkey, err := ReadRsaPubkey(filename)
	if err != nil {
		return nil, err
	}
	// 公钥进行加密
	cip, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, plainText)
	if err != nil {
		return nil, err
	}
	return cip, nil
}

//解密
func RsaPriKeyDecryptData(cipherData []byte, filename string) ([]byte, error) {
	//获取私钥进行解密
	prikey, err := ReadRsaPrikey(filename)
	if err != nil {
		return nil, err
	}
	//私钥进行解密
	plain, err := rsa.DecryptPKCS1v15(rand.Reader, prikey, cipherData)
	if err != nil {
		return nil, err
	}
	return plain, nil
}

func RsaDemoRun() {
	src := []byte("oMaoyu rsa Demo run 测试")
	cipherData, err := RsaPubKeyEncryptData(src, "./18/rsaGykey.pem")
	if err != nil {
		panic(err)
	}
	fmt.Printf("rsa 加密后数据:%x\n", cipherData)

	plainText, err := RsaPriKeyDecryptData(cipherData, "./18/rsaSykey.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(plainText))
}
