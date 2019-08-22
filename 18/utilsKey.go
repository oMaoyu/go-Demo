package _8

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

// 读取公钥
func ReadRsaPubkey(filename string) (*rsa.PublicKey, error) {
	pubKeyData,err := ioutil.ReadFile(filename)
	if err != nil {
		return nil,err
	}

	block,_ := pem.Decode(pubKeyData)

	key,err := x509.ParsePKIXPublicKey(block.Bytes)
	if public,ok := key.(*rsa.PublicKey) ;ok {
		return public,nil
	}
	return nil,err
}

// 读取私钥
func ReadRsaPrikey(filename string) (*rsa.PrivateKey, error) {
	privateData ,err := ioutil.ReadFile(filename)
	if err != nil {
		return nil,err
	}
	block , _ := pem.Decode(privateData)

	key,err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if private,ok := key.(*rsa.PrivateKey) ;ok {
		return private,nil
	}
	return nil,err
}