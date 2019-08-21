package _3

import (
	"fmt"
	"time"
)

//func Ch() {
//
//	go user1()
//	go user2()
//
//	for {
//		;
//	}
//}
//func printf(str string) {
//	for _, c := range str {
//		fmt.Printf("%c", c)
//		time.Sleep(time.Miltdsecond * 300)
//	}
//}
//func user1() {
//	printf("hahahahah")
//	ch <- 1023
//}
//func user2() {
//	<-ch
//	printf("？？？小老弟你怎么回事")
//}
//func Ch(){
//	var ch = make(chan int)
//	var ch1 = make(chan bool)
//	go func (){
//		for i := 0; i < 10; i++ {
//			ch <- i
//			fmt.Println("go 写入了",i)
//			ch1 <- true
//		}
//		close(ch)
//		close(ch1)
//	}()
//
//	//for i := 0; i < 12; i++{
//	//	if  num,ok := <- ch;ok==true {
//	//		<- ch1
//	//		fmt.Println("go 读出了",num)
//	//	}else {
//	//		fmt.Println("结束了")
//	//	}
//	//}
//
//	for num := range ch{
//		 <- ch1
//		fmt.Println(num)
//	}
//
//}

//func Ch(){
//	ch := make(chan interface{},5)
//	go scz(ch)
//	xfz(ch)
//
//	var i interface{}
//	cs(i)
//	fmt.Println(i)
//
//}
//func cs(i interface{}){
//	i = 10
//}
//
//
//func xfz(xf <-chan interface{}){
//	for num := range xf{
//		fmt.Println(num)
//	}
//}
//func scz(sc chan <-interface{}){
//	for i := 0; i < 10; i++ {
//		sc <- i*i
//	}
//	close(sc)
//}
var num = 1

func Ch() {
	ch := make(chan int)
	ch1 := make(chan bool)
	go func() {
		for i := range ch{
			num = i + 1
			fmt.Println("---子", i)
			ch1 <- true
		}
	}()
	for num < 100 {
		fmt.Println("主", num)
		ch <- num + 1
		<-ch1
	}

}


func Ch1(){
	ticker := time.NewTicker(time.Second)
	for  {
		ch := <-ticker.C
		fmt.Println(ch)
	}
}