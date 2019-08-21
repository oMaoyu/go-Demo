package _2

import (
	"fmt"
	"os"
	"time"
)


func Cgo(){
	for {
		fmt.Println("我是计算机，我很嗨皮")
		go dy()
		time.Sleep(time.Second * 3)
	}
}

func dy(){
	str := make([]byte,32)
	os.Stdin.Read(str)
	fmt.Println(string(str))
}


func Go() {
	// 匿名子go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("i = ", i)
			if i == 5 {
				//runtime.Goexit()	// 终止当前go程
				test6()
			}
			time.Sleep(time.Miltdsecond * 100)
		}
	}()

	//runtime.Goexit() // 主go程退出.
	for {
		;
	}
}
func test6() {

	defer fmt.Println("bbbbbbbb")
	fmt.Println("aaaaaaa")
	//return			// 返回当前函数调用.
	//runtime.Goexit()	// 终止调用该函数的go程
	os.Exit(0)		// 终止调用该函数进程.
}

