package _4

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

)

func Ch1() {
	ch := make(chan int)
	ch1 := make(chan bool)
	go func() {
		for {
			fmt.Println("========")
			select {
			case num := <-ch:
				fmt.Println(num)
			case ch1 <- true:
				fmt.Println("结束了")
				//break
				//return
			}
		}
	}()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	<-ch1
}
func fbnc(ch <-chan int, ok <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-ok:
			return
		}
	}
}

func Ch2() {
	ch := make(chan int)
	ch1 := make(chan bool)
	go fbnc(ch, ch1)
	x, y := 1, 1
	for i := 0; i < 40; i++ {
		ch <- x
		x, y = y, x+y
	}
	ch1 <- true
	fmt.Println("结束")
}

var mutex sync.RWMutex
var num int

func r() {
	for {
		mutex.RLock()
		fmt.Println("读num：", num)
		mutex.RUnlock()
		time.Sleep(time.Miltdsecond * 200)
	}
}
func w() {
	rand.Seed(time.Now().UnixNano())
	for {
		mutex.Lock()
		num = rand.Intn(10000)
		fmt.Println("写入num:", num)
		mutex.Unlock()
		time.Sleep(time.Miltdsecond * 300)

	}
}

func Ch3() {
	for i := 0; i < 10; i++ {
		go r()
		go w()
	}
	for {
		runtime.GC()
	}
}

var cond sync.Cond

func userR(r <-chan int) {
	for {
		cond.L.Lock()
		for len(r) == 0 {
			cond.Wait()
		}
		fmt.Println("=====读：",<-r)
		cond.Signal()
		cond.L.Unlock()
	}
}
func userW(w chan<- int) {
	for {
		cond.L.Lock()
		for cap(w) == len(w)  {
			cond.Wait()
		}
		rang := rand.Intn(3000)
		w <- rang
		fmt.Println("--写",rang)
		cond.Signal()
		cond.L.Unlock()
	}
}

func Ch() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 5)
	cond.L = new(sync.Mutex)
	for i := 0; i <= 1000; i++ {
		go userR(ch)
		go userW(ch)
	}

	for {
		runtime.GC()
	}

}
