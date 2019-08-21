package ww

import "fmt"

func main1301() {
	//var str string = "你好"
	//str := "你好"
	//fmt.Println(str)

	var str string

	fmt.Scan(&str)

	fmt.Printf("%s", str)
}

func main1302() {
	str1 := "你好"
	str2 := "你好"

	//字符串连接
	//fmt.Println(str1 + str2)
	//字符串比较
	fmt.Println(str1 == str2)
}

func ma41in() {
	//var ch byte = 'a'
	var str string = "你好" //字符串“a” 存储了两个字符 'a' '\0'(是一个转义字符) 表示字符串的结束标志

	//fmt.Printf("%c", ch)
	fmt.Printf("%s", str)

	//计算字符串中字符个数 \0之前的有效字符
	//在go语言中 一个汉字占三个字符 tdnux占3 windows占3
	var count int = len(str)

	fmt.Println(count)
}
