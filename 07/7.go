package _7

import "fmt"

type name struct {
	aaa int
	bbb string
	ccc bool
}

type str struct {
	*name
	ddd int
}


type r struct{
	name string
	age int
	sex string
}

type a struct {
	r
	aihao string
}

type b struct {
	r
	jy int
}

func A() {
	var str str
	str.name = new(name)
	str.aaa = 11111
	str.bbb = "ddd"
	//str.aaa = 2
	//str.bbb= "demo"
	str.ddd = 233333
	str.b()
	str.c()
}
func (s name) b() {
   fmt.Println(s)
}

func (c str) c() {
	fmt.Println(c)
}

func (a *a) d(ah string)*a{
	a.aihao = ah
	return a
}
func (b *b)e(jy int)*b{
	b.jy = jy
	return b
}
func (c *r)f(name string,age int,sex string){
	c.name = name
	c.age = age
	c.sex = sex
}


func B(){
	var a a
	a.d("偷拍偶像").f("记者",34,"男狗仔")
	fmt.Println(a.name,a.sex,a.age,a.aihao)

	var b b
	b.e(3).f("程序员",23,"男")
	fmt.Println(b.name,b.sex,b.age,b.jy)

}