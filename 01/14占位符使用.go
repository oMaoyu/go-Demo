package ww

import "fmt"

//int
func main1401() {
	///%% 表示字面%
	//fmt.Printf("35%%")

	a := 123456
	//%b  表示输出一个数的二进制格式(0-1)
	fmt.Printf("%b\n", a)
	//%o  表示输出一个数的八进制格式(0-7)
	fmt.Printf("%o\n", a)
	//%x %X 表示输出一个数的十六进制格式(0-9 a-f)
	fmt.Printf("%x\n", a)
	fmt.Printf("%X\n", a)

	//%d 表示输出一个数的十进制格式
	//%nd 整体数据为n位 如果不足n位在前面补空格
	//%-nd 整体数据为n位 如果不足n位在后面补空格
	//%0nd 整体数据为n位 如果不足n位在前面补0
	//如果数据超出n位 正常输出
	fmt.Printf("==%5d==\n", a)
	fmt.Printf("==%-5d==\n", a)
	fmt.Printf("==%05d==\n", a)

}

//float
func main1402() {
	a := 3.14
	//%f 是一个占位符 表示输出一个浮点型数据   默认保留六位小数 会四舍五入
	//%.nf  表示小数点后保留n位 会对第n+1位 进行四舍五入
	//%m.n m整体位数为m为 .占一位
	fmt.Printf("%f\n", a)
	fmt.Printf("%.3f\n", a)
	fmt.Printf("%-6.3f\n", a)
}

//byte 和string
func main1403() {
	var ch byte = 'a'
	fmt.Printf("%c\n", ch)
	var str string = "a"
	fmt.Printf("%s\n", str)

	//%q 是一个占位符 表示输出带双引号字符串或者字符切片
	fmt.Printf("%q\n", str)
}

//bool
func mai213n() {
	a := true
	fmt.Printf("%t\n", a)
	fmt.Printf("%T\n", a)
	//%p 是一个占位符 表示输出一个数据对应的内存地址
	fmt.Printf("%p\n", &a)

}
