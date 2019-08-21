package _5

import "fmt"

func Str(str string)map [string] int{
	dict := make(map[string]int)
	str1 := []rune(str)
	for _,j:= range str1 {
		dict[string(j)] ++
	}
	return dict
}

type Name struct {
	num int
	name string
	Sdemo demo
 }
type demo struct {
	num int
	name string
}

func StructDemo(){
	name := Name{10,"haha",demo{1,"23333"}}
	fmt.Println(name)
}


//func Demo(){
//	var arry []*int
//	i := 10
//	j := &i
//	k := 20
//	fmt.Println(j)
//	arry = append(arry, j)
//	arry = append(arry,&k)
//	fmt.Println(arry)
//}