package _8
/*




 */
//import "C"
import (
	"fmt"
)

type i struct {
	n1 int
	n2 int
}
type jian struct {
	i
}
type jia struct {
	i
}

type j interface {
	j() int
}


func (j *jian) j()int{
	return j.n1 - j.n2
}

func (j *jia) j()int{
   return j.n1 + j.n2
}


func face(face j){
	fmt.Println(face.j())
}


func A(){
	f := []interface{}{"dfa"}
	if  jj,ok:=f[0].(string);ok{
		fmt.Println(jj)
	}


    j1 := jia{i{10,31}}
	j2 := jian{i{81,26}}
    face(&j1)
	face(&j2)

	i1 := dog{"妞妞"}
	i2 := roo{"壮壮"}
	eat(&i1)
	eat(&i2)
	run(&i1)
	run(&i2)

}

type Animal interface {
	eat()
	run()
}

type dog struct {
	name string
}
type roo struct {
	name string
}
func eat(face Animal){
	face.eat()
}
func run(face Animal){
	face.run()
}
func (d *dog)eat(){
	fmt.Println("dog,eat",d.name)
}
func (d *dog)run(){
	fmt.Println("dog,run",d.name)

}
func (d *roo)eat(){
	fmt.Println("roo,eat",d.name)

}
func (d *roo)run(){
	fmt.Println("roo,run",d.name)
}
