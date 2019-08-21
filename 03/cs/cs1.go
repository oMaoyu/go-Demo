package cs
// num  弹跳次数
//c 函数参数名
func Dg(num int) (c func(num float64) (float64, float64)) {
	//i 弹跳次数  j  总高度累计
	i, j := num, 0.0
	// num 下一次弹跳高度  返回值1 弹跳总高度 返回2 在次弹跳高度
	return func(num float64) (f float64, f2 float64) {
		i-- // 次数-1
		j += num + num / 2.0 // 累计高度
		if i == 0 {
			return j - num / 2.0, num / 2.0
		} else {
			return c(num / 2.0) //传入下次弹跳高度
		}
	}
}
