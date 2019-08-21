package ww

import "fmt"

func mai111n() {
	//单精度浮点型  小数位数有效为7位
	var a float32 = 1.234
	//双精度浮点型  小数位数有效为15位
	var b float64 = 1.234
	//建议在开发中使用float64代替float32

	//自动推导类型创建浮点型 默认为float64
	c := 3.14//float64

	//fmt.Println(a)
	//fmt.Println(b)

	fmt.Printf("%T\n", c)

	fmt.Printf("%.6f\n", a)
	fmt.Printf("%.6f\n", b)
}
