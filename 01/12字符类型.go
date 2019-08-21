package ww

import "fmt"

func main1201() {
	//var 变量名 byte = ''

	//每一个字符对应都有一个ASCII值
	//var ch byte = 'a'
	//var ch byte = 'A'

	//转义字符
	var ch byte = '\v'

	//在println打印时 打印字符对应的 十进制整型的值
	fmt.Println(ch)

	//%c  是一个占位符 表示输出一个字符
	//fmt.Printf("%c", ch+32)
}

func mai11211n() {
	//var a uint8 ='a'

	var a byte
	//fmt.Scan(&a)

	//如果通过键盘接收字符  必须使用fmt.Scanf
	fmt.Scanf("%c", &a)
	fmt.Printf("%c", a)
}
