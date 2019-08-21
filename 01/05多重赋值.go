package ww

import "fmt"

func main0501() {

	//多重赋值
	//定义多个变量
	//var a, b, c int = 10, 20, 30

	//自动推导类型创建多个变量
	//a, b, c := 10, 20, 30
	a, b, c := 10, 3.14, "hello"

	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	//\n 转义字符  表示换行
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

func mai1111n() {
	//匿名变量
	//_匿名变量  在后面学习中涉及到函数多个返回值时 使用匿名变量
	a, _, _ := 10, 20, 30

	fmt.Println(a)
	//fmt.Println(_)//err 匿名变量数据不能打印
	//fmt.Println(b)
}
