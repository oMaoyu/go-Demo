package _6

import "fmt"

type Person struct {
	str string
	is bool
	num int
	ary []string
}


func A()  {
	var p Person
	initPerson(&p)
	fmt.Println(p)
}

func initPerson(p *Person)  {
	p.ary = []string{"hahha","hehehe"}
	p.is = true
	p.num = 11
	p.str = "???"
}
