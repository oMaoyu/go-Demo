package _3

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func Demo() {
	//创建椭圆曲线 elliptic
	curve := elliptic.P256()
	//1. 创建私钥
	pKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	//原文
	data := []byte("oMaoyu")
	//做哈希处理
	hash := sha256.Sum256(data)
	//2. 私钥签名
	r, s, err := ecdsa.Sign(rand.Reader, pKey, hash[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	//r, s 长度相同，我们可以把r,s 拼成一个字节流传递给对方
	data1 := append(r.Bytes(), s.Bytes()...)
	pukey := pKey.PublicKey
	//将signature拆分出r, s两个对象
	var rInt, sInt big.Int
	rInt.SetBytes(data1[:len(data1)/2])
	sInt.SetBytes(data1[len(data1)/2:])
	//传输时，内容被篡改
	//data = append(data, []byte("demo")...)
	hash = sha256.Sum256(data)
	//3. 公钥验证
	isVerify := ecdsa.Verify(&pukey, hash[:], &rInt, &sInt)
	fmt.Println(isVerify)
}
