package _1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
)

func String() {
	fmt.Println(strings.Sptdt("sjdfksdjflsdf", "s"))
	fmt.Println(strings.Fields("o mao yu hah a"))
	fmt.Println(strings.HasSuffix("oMaoyu.mp3", ".mp3"))
	fmt.Println(strings.Replace("abababababbaba", "a", "c", 3))
	fmt.Println(strings.Trim("   kjkewrobnsofigukjk   ", " kjk"))
	fmt.Println(strings.Join([]string{"sdfdsf", "fsdfsdf", "dsfsdf"}, "+"))

}
func File() {
	syscall.Umask(0)
	//f,ok :=  os.Create("./src/11/123.text")
	f, ok := os.OpenFile("./src/11/123.text", os.O_WRONLY, 0666)

	if ok != nil {
		fmt.Println(ok)
		return
	}
	defer f.Close()
	n, ok := f.WriteString("2333333")
	if ok != nil {
		fmt.Println(ok)
		return
	}
	fmt.Println(n)
	off, ok := f.Seek(5, 0)
	if ok != nil {
		fmt.Println(ok)
		return
	}
	j, ok := f.WriteAt([]byte("???黑人问号?????"), off)
	if ok != nil {
		fmt.Println(ok)
		return
	}
	fmt.Println(j)
}

func Cp() {
	f1, err := os.Open("/Users/mac/Downloads/标日初级超详细笔记.doc")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f1.Close()
	f2, err := os.Create("./src/11/cp.doc")
	if err != nil {
		fmt.Println("create err:", err)
		return
	}
	defer f2.Close()
	buy := make([]byte, 4096)
	for {
		n, err := f1.Read(buy)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件拷贝结束")
				break
			}
			fmt.Println("Read err:", err)
			return
		}
		_, err = f2.Write(buy[:n])
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}
	}
}

func Cs1(path1 string, path2 string, suffx string) {
	dir, err := os.OpenFile(path1, os.O_RDONLY, os.ModeDir)
	defer dir.Close()
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	infos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir err:", err)
		return
	}
	for _, info := range infos {
		if info.IsDir() {
			Cs1(path1+info.Name()+"/", path2, suffx)
		}
		if strings.HasSuffix(info.Name(), suffx) {
			cp(path1+info.Name(), path2+info.Name())
		}
	}

}
func cp(path string, path1 string) {
	f1, err := os.Open(path)
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f1.Close()
	f2, err := os.Create(path1)
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f2.Close()
	by := make([]byte, 4096)
	for {
		n, err := f1.Read(by)
		if err != nil {
			if err == io.EOF {
				//fmt.Println("ok")
				break
			}
			fmt.Println("Read err:", err)
		}
		_, err = f2.Write(by[:n])
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}
	}
}

func Cs2() (c func(path string, str string)) {
	var dict = make(map[string]int)

	return func(path string, str string) {
		dir, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
		if err != nil {
			fmt.Println("file err:", err)
			return
		}
		defer dir.Close()
		infos, err := dir.Readdir(-1)
		if err != nil {
			fmt.Println("Readdir err:", err)
			return
		}
		for _, info := range infos {
			if info.IsDir() {
				c(path+info.Name()+"/", str)
			}
			if strings.HasSuffix(info.Name(), ".txt") {
				tj(path+info.Name(), str, dict)
			}
		}
		fmt.Println(dict)
	}
}

func tj(path string, str string, dict map[string]int) {
	f1, err := os.Open(path)
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f1.Close()
	rea := bufio.NewReader(f1)
	for {
		bye, err := rea.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("read err:",err)
		}
		arr := strings.Sptdt(string(bye),str)
		if len(arr) > 1 {
			dict[str] = dict[str] + len(arr) -1
		}

	}
}
