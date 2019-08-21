package ww

import "fmt"

func main0401() {

	//var a int = 10
	//自动推导类型(根据值来确定变量的类型)
	//变量名:=值
	a := 10   //int
	b := 3.14 //float64

	//fmt.Println(a)
	//fmt.Println(b)

	//通过格式化打印数据类型
	//%T 是一个占位符 表示输出一个数据对应的类型
	fmt.Printf("%T", a)
	fmt.Println()
	fmt.Printf("%T", b)

}

//func mai111n() {
//
//	//var a string = "hello"
//	//通过自动推导类型创建字符串
//	a := "你好" //创建变量
//	a = "我不好" //赋值
//
//	fmt.Println(a)
//	fmt.Printf("%T", a)
//}
