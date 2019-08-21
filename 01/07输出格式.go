package ww

import "fmt"

func main0701() {
	//打印输入并换行
	//fmt.Println("hello")
	//打印不换行
	//\n 是一个转义字符 表示换行 是一个控制字符
	fmt.Print("hello\n")
	fmt.Print("world")



}

//func main111() {
//	a := 10
//	b := 3.1415
//	c := "你瞅啥"
//	//fmt.Printf("%T\n",a)
//	//%d 是一个占位符 表示输出一个十进制整型数据
//	fmt.Printf("%d\n", a)
//	//%f 是一个占位符 表示输出一个浮点型数据   默认保留六位小数 会四舍五入
//	//%.3f  表示小数点后保留三位 会对第四位进行四舍五入
//	fmt.Printf("%.3f\n", b)
//	//%s 是一个占位符 表示输入一个字符串类型数据
//	fmt.Printf("%s\n",c)
//}
