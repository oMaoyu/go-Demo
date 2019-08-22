package _8

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// 私钥签名
func rsaPriKeySignData(data []byte, filename string) ([]byte, error) {
	// 首先获取私钥
	prikey, err := ReadRsaPrikey(filename)
	if err != nil {
		return nil, err
	}
	// 对原始数据进行哈希计算
	s256 := sha256.Sum256(data)
	// 开始签名
	// 随机数生成器   私钥  加密方法    哈希切片
	sign, err := rsa.SignPKCS1v15(rand.Reader, prikey, crypto.SHA256, s256[:])
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// 公钥验证
func rsaPubKeyVerifyData(data []byte, filename string, signData []byte) bool {
	// 获取公钥
	pubkey, err := ReadRsaPubkey(filename)
	if err != nil {
		return false
	}
	// 获取原始数据哈希值

	s256 := sha256.Sum256(data)
	// 开始验证签名
	//func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error {
	err = rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, s256[:], signData)
	if err != nil {
		return false
	}
	return true
}

func RsaSignatureDemo(){
	data := []byte("随随便便写了一个数据,这个数据即将用来进行签名验证")

	signData ,err := rsaPriKeySignData(data,"./18/rsaSyKey.pem")
	if err != nil {
		panic(err)
	}

	isSign := rsaPubKeyVerifyData(data,"./18/rsaGyKey.pem",signData)
	fmt.Println(isSign)
}