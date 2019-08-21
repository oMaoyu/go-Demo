package ww

import "fmt"

func main0801() {
	var a int
	//& 取地址运算符  引用运算符
	//fmt.Scan是一个阻塞式的请求
	fmt.Scan(&a)
	// 取出变量对应的内存地址
	//fmt.Println(&a)

	fmt.Println(a)
}

func main0802() {
	//多个值输入
	var a float64
	var b float64
	//如果输入有多个值 可以以换行和空格作为输入的结束
	fmt.Scan(&a, &b)

	fmt.Println(a, b)
}

//func mai111n() {
//	//格式化输入
//	//var a int
//	//var b float64
//	var __ int =456
//	//fmt.Scanf("%d", &a)
//	//fmt.Scanf("%f", &b)
//
//	//fmt.Println(a)
//	//fmt.Println(b)
//	fmt.Println(__)
//}
