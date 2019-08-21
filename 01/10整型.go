package ww

import "fmt"

func main1001() {
	//int uint (unsigned int)

	var a int = -10

	//uint  表示无符号整型数据 只能存储大于等于0的数
	var b uint = 10

	fmt.Println(a)
	fmt.Println(b)

}

func main21() {
	//数据溢出 （超出数据取值范围）
	//如果数据到达了区间的最值时在操作就会发生变化

	var a int16 = 32767
	a = a + 1

	fmt.Println(a)
}
