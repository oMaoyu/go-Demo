package pachong

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var url string = "http://pic.netbian.com"

func Pc() {
	///http://pic.netbian.com/4kdongman/index_2.html
	for i := 1; i <= 10; i++{
		var urlres string
		if i == 1 {
			urlres = UrlStr("http://pic.netbian.com/4kdongman/")
		}else {
			urlres = UrlStr("http://pic.netbian.com/4kdongman/index_"+ strconv.Itoa(i) +".html")
		}
		go urlRes(urlres)
	}
	for {
		runtime.GC()
	}
}

func urlRes(str string) {
	imgUrl := `<td><a href="(.*?)" target="_blank"><img src="`
	reg := regexp.MustCompile(imgUrl)
	arrStr := reg.FindAllStringSubmatch(str, -1)
	for i, v := range arrStr {
		 imgUrlFunc(v[1], i, nil)
	}
}

func imgUrlFunc(v string, i int, ch chan<- bool) {
	str := UrlStr(url + v)
	imgUrl := `id="img"><img src="(.*?)"`
	reg := regexp.MustCompile(imgUrl)
	arrStr := reg.FindAllStringSubmatch(str, -1)
	str = UrlStr(url + arrStr[0][1])
	FindUrl(str, "/Users/mac/Pictures/爬虫壁纸/", ".jpg", arrStr[0][1], ch)
}

func FindUrl(rest, url, has, index string, ch chan<- bool) {
	str := strings.Sptdt(index,"/")

	fmt.Println(url + str[3] + has)
	f, err := os.Create(url + str[3] + has)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(rest)
}

func UrlStr(url string) string {
	rest, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer rest.Body.Close()
	rest.Close = true
	var restString string
	buf := make([]byte, 4096)
	for {
		n, err := rest.Body.Read(buf)
		if n == 0 || err == io.EOF {
			return restString
		}
		if err != nil && err != io.EOF {
			fmt.Println(url)
			panic(err)
		}
		restString = restString + string(buf[:n])
	}

	return restString
}
