package data

import (
	"fmt"
	"math/rand"
	"time"
)

// 冒泡排序
func bubbling(arr []int) {
	flg := false
	for i := 0; i < len(arr)-1; i++ {
		if flg {
			return
		}
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 选择
func Select(arr []int) {
	start := 0
	for i := 0; i < len(arr)-1; i++ {
		index := 0
		for j := 1; j < len(arr)-i; j++ {
			start++
			if arr[index] < arr[j] {
				index = j
			}
		}
		arr[len(arr)-i-1], arr[index] = arr[index], arr[len(arr)-i-1]
	}
	fmt.Println(start)
}

//  插入
func interposition(arr []int) {
	start := 0
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			start++
			if arr [j] > arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	fmt.Println(start)
}

// 计数统计排序  用map 大数据最快

// 希尔排序
func shellSort(arr []int) {
	start := 0
	for logo := len(arr) / 2; logo > 0; logo = logo / 2 {
		for i := logo; i < len(arr); i++ {
			for j := i - logo; j >= 0; j = j - logo {
				start++
				if arr[j] > arr[j+logo] {
					arr[j], arr[j+logo] = arr[j+logo], arr[j]
				} else {
					break
				}
			}
		}
	}
	fmt.Println(start)
}

//快速排序
func quicksort(arr []int, l, r int) []int {
	if l < r {
		i, j, data := l, r, arr[l]
		for i < j {
			// 必须先从右边换好标签后才能开始左边前行
			for i < j && arr[j] <= data {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
			for i < j && arr[i] >= data {
				i++
			}
			arr[i], arr[j] = arr[j], arr[i]

		}
		quicksort(arr,l,i-1)
		quicksort(arr,i+1,r)
	}

	return arr
}

func NewSort() {
	arr := make([]int, 0)
	rand.Seed(time.Now().UnixNano()) // 播种子
	// 打造 十万个 在 0-999 之间的 数据集
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(5000))
	}
	//bubbling(arr)
	//Select(arr)
	//interposition(arr)
	//shellSort(arr)
	quicksort(arr,0, len(arr)-1)
	fmt.Println(arr)
}
