package _8

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateKeyPair(bits int) error {

	//1. 创建私钥 (rsa包)
	syKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	//2. 对私钥进行编码，得到der格式的字符串（秘钥相关的格式）
	der, err := x509.MarshalPKCS8PrivateKey(syKey)
	if err != nil {
		return err
	}
	//3. 将der字符串拼装到pem格式的数据块中
	pemBlock := pem.Block{
		Type:    "My oMaoyu RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   der,
	}
	//4. 对pem格式进行base64编码，最终得到私钥数据，写入文件
	f1, _ := os.Create("./18/rsaSyKey.pem")
	defer f1.Close()
	err = pem.Encode(f1, &pemBlock)
	if err != nil {
		return err
	}

	//生成der格式数据
	publicKey := syKey.PublicKey
	der, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	//拼接pemblock
	pemBlock = pem.Block{
		Type:    "My oMaoyu RSA PUBLICK KEY",
		Headers: nil,
		Bytes:   der,
	}
	//写入文件中
	f2, _ := os.Create("./18/rsaGyKey.pem")
	defer f2.Close()
	err = pem.Encode(f2, &pemBlock)
	if err != nil {
		return err
	}

	return nil
}
