package main

import "fmt"

func D()  {
	j := demo(10)
	fmt.Println(j(1,2))
	fmt.Println(j(1,2))
	fmt.Println(j(1,2))
	fmt.Println(j(1,2))
	fmt.Println(j(1,2))
}

func demo(i int) (c func(j int, k int) int) {
	g := i
	return func(j int, k int) int {
		g = j + k + g
		return g
	}
}

func add(arr ...int) {

}
