package main

import _8 "demo/18"

type Func func(int, int) int

func main() {
	//app := iris.New()
	//
	//app.Get("/", func(iris.Context) {
	//	fmt.Print("???")
	//})
	//app.Run(iris.Addr(":8080"))

	//var f Func
	//f = func(i int, i2 int) int {
	//	return i + i2
	//}
	//fmt.Println(f(41,23))
	//
	//D()
	//ss.DemoD(1,2)
	//num 弹跳次数
	//v := cs.Dg(10)
	//
	//i, j := v(100.0)
	//
	//fmt.Printf("总共跳了%.2f米，最后一次弹%.2f米\n", i, j)

	//_4.T()
	//_4.Mp()
	//_4.Xz()
	//_1.Cp()
	//_4.Sj()
	//_2.Go()
	//_4.Q()
	//dict := _5.Str("aaabbeeqqqq")
	//fmt.Println(dict)
	//
	//_5.StructDemo()
	//
	////_5.Demo()
	//_6.A()
	//_7.A()
	//_7.B()
	//_8.A()

	//fmt.Println(_0.Str("adfadfFDAKDFD"))
	//_1.String()
	//_1.File()
	//_1.Cp()
	//_1.Cs1("./src/","/Users/mac/Music/QQ音乐/",".go")
	//f := _1.Cs2()
	//f("/Users/mac/Music/QQ音乐/cs/","oMaoyu")
	//_2.Go()
	//_2.Cgo()
	//_4.Ch()
	//_3.Ch1()

	//ch := make(chan bool)
	//
	//go func (){
	//	for i := 0; i < 5; i++ {
	//		ch <- true
	//		time.Sleep(time.Second)
	//	}
	//}()
	//
	//for _ = range ch {
	//	fmt.Println("？？？")
	//}

	//pachong.Pc()
	//_6.M()

	//data.NewArr()
	//data.NewNode()
	//data.NewTwoNode()
	//data.NewLoopNode()
	//data.NewStackNode()
	//data.NewBinary()
	//data.NewSort()
	//data.MpDemo()

	//data.NewHash()
	//data.SetKey("1312", "dfad")
	//data.SetKey("世界和平", "核平")
	//data.SetKey("go", "宝可梦")
	//data.SetKey("hahah", "哈哈哈")
	//data.SetKey("???", 123213)
	//data.SetKey("欸", "?")
	//data.SetKey("go", "狗朗")
	//fmt.Println(data.GetKey("???"))
	//for i := 0; i < len(data.Arr); i++ {
	//	data.Arr[i].Log()
	//	fmt.Println("长度:", data.Arr[i].Length())
	//}

	//key := []byte("sdfgqwer12345678")
	//src := []byte("aes对称加密的第一个尝试,oMaoyu")
	//text, err := _7.AesCtrEncrypt(key, src)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//text, _ = _7.AesCtrDecrypt(key, text)
	//fmt.Println(string(text))


	//key := []byte("sdfgqwer")
	//src := []byte("des对称加密的第一个尝试,oMaoyu")
	//text, err := _7.DesCBCEncrypt(key, src)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//text, _ = _7.DesCBCDecrypt(key, text)
	//fmt.Println(string(text))
	_8.GenerateKeyPair(1024)
}

func ddd(dict map[string]int) {
	dict["o"]++
}
