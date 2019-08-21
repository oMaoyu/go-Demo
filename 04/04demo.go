package _4

import (
	"fmt"
	"math/rand"
	"time"
)

func T() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	fmt.Println(arr)
}

func Mp() {
	arr := [10]int{71, 12, 33, 52, 23, 91, 0, 11, 9, 10}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

//选择排序
func Xz() {
	arr := [10]int{71, 12, 33, 52, 23, 91, 0, 11, 9, 10}
	k := 0
	for i := 0; i < len(arr)-1; i++ {
		k = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[k] {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
	fmt.Println(arr)
}

//创建随机数
func Sj() {
	f, i := array1(), 0
	arr := [10]int{0}
	arr1:=[10]string{"银灰","斯卡蒂","夜莺","白面鸮","法华玲","蓝毒","空","初雪","守林","陨星"}
	for i != cap(arr) {
		arr[i] = f(cap(arr))
		i++
	}
	for _,v := range arr{
		fmt.Println(arr1[v])
	}
}

// 直接返回无重复值的随机数
func array1() (f func(i int) int) {
	var arr = [100]int{0}
	rand.Seed(time.Now().UnixNano())
	return func(i int) int {
		j := rand.Intn(i)
		if arr[j] == 0 {
			arr[j] = 1
			return j
		} else {
			//fmt.Println("出现重复数", j)
		}
		return f(i)
	}
}

// 判断是否有重复数 返回bool值
func array() (func(i int) bool) {
	var arr = [100]int{0}
	return func(i int) bool {
		if (arr[i] == 0) {
			return true
		} else {
			arr[i] = 1
			return false
		}
	}
}

//============================================================》切片

func Q() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[4:]
	s[0] = 88
	fmt.Println(s, arr)
	//fmt.Printf("%p,%p", s, arr)

}
