package ww

import "fmt" // format 包 格式化输入输出

func i() {
	fmt.Println("")

	//	%T 占位符 打印数据类型

}

//func main111() {
//	i, j := yuan(4.8)
//	i, j = j, i
//
//	fmt.Println(i, j)
//
//}

// 圆面积 周长
func yuan(r float64) (float64, float64) {
	var PI = 3.1415926
	return PI * r * r, PI * r
}
