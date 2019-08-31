package _1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

// 定义钱包  包含公钥私钥
type wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte // 存放拼接的字节流  可以拆分出公钥
}

// 初始化钱包
func newWallet() *wallet {
	pKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	puKey := pKey.PublicKey
	puKeyByte := append(puKey.X.Bytes(), puKey.Y.Bytes()...)

	return &wallet{pKey, puKeyByte}

}

// 生成地址
func (w *wallet) getAdd()string {
	// 根据比特币的公钥算法 先进行256加密 在进行160加密
	hash := sha256.Sum256(w.PublicKey)
	h := ripemd160.New()
	h.Write(hash[:])
	pubKeyHash := h.Sum(nil)
	// 然后拼接一个字节 形成21字节
	payload := append([]byte{byte(00)},pubKeyHash...)
	//求checksum，获取4字节数据
	hash = sha256.Sum256(payload)
	hash = sha256.Sum256(hash[:])
	checksum := hash[0:4]
	payload = append(payload,checksum...)
	return base58.Encode(payload)
}
