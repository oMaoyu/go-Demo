package _7

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
)

//des:
//秘钥：8字节
//分组长度: 8字节

//cbc:
// 需要填充，长度与des一致，8字节

//en前缀:表示肯定
//de，un前缀：表示否定

//加密函数
func DesCBCEncrypt(key, plainText []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	byte := bytes.Repeat([]byte("1"), block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, byte)

	plainText,_ = paddingNumber(plainText,blockMode.BlockSize())

	blockMode.CryptBlocks(plainText,plainText)

	return plainText,err
}

//填充函数
//参数1：明文数据
//参数2：每一个分组的长度
func paddingNumber(src []byte, blocksize int) ([]byte, error) {
	//校验src不能为nil
	if src == nil {
		return nil, errors.New("填充数据时，原内容不应为nil")
	}

	fill := len(src) % blocksize

	fill = blocksize - fill

	b := byte(fill)

	newByte := bytes.Repeat([]byte{b}, fill)

	src = append(src, newByte...)

	return src, nil
}

//去掉填充的函数
func unPaddingNumber(src []byte) []byte {
	l := int(src[len(src)-1])

	return src[:len(src)-l]
}

//解密函数
func DesCBCDecrypt(key, cipherText []byte) ([]byte, error) {
	fmt.Println("开始解密....")
	//第一步：创建密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := bytes.Repeat([]byte("1"), block.BlockSize())

	//第二步：创建CBC分组
	mode := cipher.NewCBCDecrypter(block, iv)

	//第三步：解密
	mode.CryptBlocks(cipherText /*明文*/, cipherText /*密文*/)
	//在解密之后进行数据截取
	cipherText = unPaddingNumber(cipherText)

	return cipherText /*明文*/, nil
	//return []byte("解密后的数据"), nil
}
