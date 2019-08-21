package _6

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

//const oMaoyu = "oMaoyu"
const oMaoyu = "json"

type user struct {
	Name   string `oMaoyu:"-"isBool:"true"`
	age    int    `oMaoyu:"Age.hahah"`
	text   string `oMaoyu:"???"`
	isBool bool   `oMaoyu:"sss"`
	hahaha string
}

type User1 struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Msg    string `json:"msg"`
	BgImg  string `json:"bg_img"`
	IsBool bool   `json:"is_bool"`
}

func (u user) Add() {
	//fmt.Println(u.age)
}
func (u user) Btc(str1, str2 string) string {
	return str1 + str2 + u.Name
}
func (u user) Aaa(str1, str2, str3, str4 string) (string, int, bool) {
	return u.Name, u.age, u.isBool
}

//func reInter(i interface{}) {
//	reVal := reflect.ValueOf(i)
//	reType := reflect.TypeOf(i)
//	if reVal.Kind() == reflect.Ptr && reVal.Elem().Kind() == reflect.String {
//		reVal.Elem().SetString("hehehe")
//	}
//	inter := reVal.Interface()
//	if n, ok := inter.(int); ok {
//		fmt.Println(n + 10)
//	}
//	if n, ok := inter.(string); ok {
//		fmt.Println(n + "hahah")
//	}
//	if reType.Kind() == reflect.Struct {
//		for i := 0; i < reType.NumField(); i++ {
//			field := reType.Field(i)
//			val := reVal.Field(i)
//			fmt.Println(field.Name, field.Type, field.Tag.Get(oMaoyu), val)
//		}
//	}
//}
func M() {
	//user := user{"oMaoyu", 21, "ooooo", true, "JFSF"}
	//reInter(&user)
	//reFunc(&user)

	user := User1{"oMaoyu", 18, "千胜白指日可待", "www.baidu.com",false}
	reJson(&user)
	buf, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))

}
func reInter(i interface{}) {
	reVal := reflect.ValueOf(i)
	if reVal.Kind() == reflect.Ptr && reVal.Elem().Kind() == reflect.Struct {
		reVal = reVal.Elem()
		for i := 0; i < reVal.NumField(); i++ {
			tField := reVal.Type().Field(i)
			name := ""
			val := ""
			if str := tField.Tag.Get(oMaoyu); str != "" {
				name = str
			} else {
				name = tField.Name
			}
			if field := reVal.Field(i); field.Kind() == reflect.String {
				if str := tField.Tag.Get("isBool"); str != "" {
					val = strings.ToTitle(field.String())
				} else {
					val = strings.ToLower(field.String())
				}
			} else {
				val = fmt.Sprint(reVal.Field(i))
			}

			fmt.Println(name, val)
		}
	}
}
func reFunc(i interface{}) {
	reVal := reflect.ValueOf(i)
	if reVal.Kind() == reflect.Ptr && reVal.Elem().Kind() == reflect.Struct {
		reVal = reVal.Elem()

		var arr []reflect.Value
		for i := 0; i < 4; i++ {
			arr = append(arr, reflect.ValueOf("11"))
		}
		reVal.Method(1).Call(nil)
		arr1 := reVal.Method(0).Call(arr)
		fmt.Println(arr1[2])

	}
}
func reJson(i interface{}) {
	reVal := reflect.ValueOf(i)
	if reVal.Kind() == reflect.Ptr && reVal.Elem().Kind() == reflect.Struct {
		reVal = reVal.Elem()
		Str := "{"
		for i := 0; i < reVal.NumField(); i++ {
			tField := reVal.Type().Field(i)
			name := ""
			val := ""
			if str := tField.Tag.Get(oMaoyu); str != "" {
				name = str
			} else {
				name = tField.Name
			}
			val = fmt.Sprint(reVal.Field(i))
			//{"name":"oMaoyu","age":18,"msg":"千胜白指日可待","bg_img":"www.baidu.com"}
			sss := ""
			if reVal.Field(i).Kind() == reflect.String {
				sss = "\""
			}

			if i != reVal.NumField()-1 {
				Str = Str + "\"" + name + "\":" + sss + val + sss + ","
			} else {
				Str = Str + "\"" + name + "\":" + sss + val + sss + "}"
			}

		}
		fmt.Println(Str)
	}
}
