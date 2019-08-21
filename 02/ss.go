package ss

import "fmt"

//func main() {
//	i := 11
//	demoB(i, i*2+1)
//	demoC(66, 139432, 4, 1)
//	demoD(192384, 16)
//	demoD1(192384, 16)
//}
func DemoD(a int, b int) {
	j, k, u := a, b, 0
	for {
		//过程
		//fmt.Println(j,k)
		if j == 0 || k == 0 {
			fmt.Println("最大公因数是：", j+k, "最小公倍数是：", a*b/(j+k))
			fmt.Println(u)
			break
		}
		if j > k {
			j = j % k
		} else if j < k {
			k = k % j
		}
		u++
	}
}
func DemoD1(a int, b int) {
	j, k, u := a, b, 0
	for {
		//过程
		//fmt.Println(j,k)
		if j == k {
			fmt.Println("最大公因数是：", k, "最小公倍数是：", a*b/k)
			fmt.Println(u)
			break
		}
		if j > k {
			j = j - k
		} else if j < k {
			k = k - j
		}
		u++
	}
}

func demoC(num int, s int, y int, r int) {
	fmt.Println(num/7, "周", num%7, "天：")
	fmt.Println(s/(3600*24), "天", s/3600%24, "时", s/60%60, "分", s%60, "秒")
	j, t := y, r
	for {
		j--
		if j == 0 {
			break
		}
		switch j {
		case 1, 3, 5, 7, 8, 10, 12:
			t += 31
		case 4, 6, 9, 11:
			t += 30
		case 2:
			t += 28
		}
	}
	fmt.Println("是这年的第：", t, "天")
}

func demoB(n int, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m/2-i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < i*2+1; j++ {
			if i == n-1 || j == 0 || j == i*2 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func demoA() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a > b && a > c {
		fmt.Println("a重")
	} else if b > a && b > c {
		fmt.Println("b重")
	} else if a == b && b == c {
		fmt.Println("三头猪都重")
	} else {
		fmt.Println("c重")
	}
}
