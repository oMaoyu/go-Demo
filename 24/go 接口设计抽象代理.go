package _4

import "fmt"

// 接口抽象测试
// 组装电脑
type Computer struct {
	cpu Cpu
	car Card
	mem Memeory
}
type Cpu interface {
	calculate()
}

type Card interface {
	display()
}

type Memeory interface {
	storage()
}

func NewComputer(cpu Cpu, car Card, mem Memeory) *Computer {
	com := Computer{cpu:cpu,car:car,mem:mem}
	return &com
}

// 对应产品类

type ICpu struct {}

func (c *ICpu)calculate()  {
	fmt.Println("因特尔的cpu运行了...")
}
type ICard struct {}
type NCard struct {}

func (c *ICard)display()  {
	fmt.Println("因特尔的显卡运行了...")
}
func (c *NCard)display()  {
	fmt.Println("NVIDIA 的显卡运行了...")
}
type IMem struct {}
type KMem struct {}

func (m *IMem)storage()  {
	fmt.Println("因特尔的内存运行了...")
}

func (m *KMem)storage()  {
	fmt.Println("King的内存运行了...")
}
func(com *Computer) run(){
	com.cpu.calculate()
	com.car.display()
	com.mem.storage()
}


func DemoRun() {
	NewComputer(&ICpu{},&ICard{},&IMem{}).run()
	NewComputer(&ICpu{},&NCard{},&KMem{}).run()
}
