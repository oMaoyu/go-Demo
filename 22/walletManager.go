package _1

import (
	"errors"
)

// 负责管理钱包
//	存储钱包公私钥 以及转换成地址

type walletManager struct {
	wallets map[string]*wallet
}

// new一个钱包
func newWalletManager() *walletManager {
	var wm walletManager
	wm.wallets = make(map[string]*wallet)
	return &wm
}

//创建一个新的钱包,返回对应的地址
func (wm *walletManager) createWallet() (string, error) {
	w := newWallet()
	if w == nil {
		return "", errors.New("创建失败")
	}
	add := w.getAdd()
	return add,nil
}

// 存储钱包到本地
