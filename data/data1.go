package data

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type arr struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
// 设置每个元素的内存地址大小
const Tag = 8
// 初始化
func (a *arr) New(l, c int, data ...int) {
	if a == nil {
		return
	}
	if data == nil {
		return
	}
	if l < 0 || c < 0 || l > c || len(data) > l {
		return
	}
	a.Data = C.malloc(C.size_t(c) * Tag)
	a.Len = l
	a.Cap = c

	p := uintptr(a.Data)

	for i, v := range data {
		*(*int)(unsafe.Pointer(p + uintptr(i*Tag))) = v
	}
}
// 打印
func (a *arr) Printf() {

	if a == nil || a.Data == nil {
		return
	}
	p := uintptr(a.Data)
	for i := 0; i < a.Len; i++ {
		fmt.Printf("%d  ", *(*int)(unsafe.Pointer(p + uintptr(i*Tag))))
	}
	fmt.Println()
}
// 添加
func (a *arr) Add(data ...int) {
	if a == nil || a.Data == nil {
		return
	}
	for a.Cap < a.Len+len(data) {
		a.Dilatation()
	}
	p := uintptr(a.Data)
	p += Tag * uintptr(a.Len)
	for i, v := range data {
		*(*int)(unsafe.Pointer(p + uintptr(i*Tag))) = v
	}
	a.Len += len(data)
}
// 扩容内存
func (a *arr) Dilatation() {
	a.Data = C.realloc(a.Data, C.size_t(a.Cap)*2*Tag)
	a.Cap *= 2
}
// 获取对应下标的int指针
func (a *arr)referIntPrt(i int)*int  {
	err := -1
	if a.Data == nil || a == nil {
		return &err
	}
	if i < 0 || i >= a.Len {
		return &err
	}
	p := uintptr(a.Data)
	p += uintptr(i) * Tag
	return (*int)(unsafe.Pointer(p))
}
// 查询
func (a *arr) ReferIndex(i int) int {


	return *a.referIntPrt(i)
}
// 修改
func (a *arr)Alter(index ,value int){
	*a.referIntPrt(index) = value
}
// 查询对应值的下标
func (a *arr) ReferValue(value int) int {
	if a.Data == nil || a == nil {
		return -1
	}
	p := uintptr(a.Data)

	for i := 0; i < a.Len; i++ {
		v := *(*int)(unsafe.Pointer(p + uintptr(i)*Tag))
		if v == value {
			return i
		}
	}
	return -1
}
//指定位置之后添加  or 插入
func (a *arr) AddIndex(index int, data ...int) {
	if a.Data == nil || a == nil {
		return
	}
	if index < 0 {
		return
	}
	if index >= a.Len-1 {
		for _, v := range data {
			a.Add(v)
		}
		return
	}
	for a.Len+len(data) > a.Cap {
		a.Dilatation()
	}
	// 获取本次添加后最后元素的指针地址
	p := uintptr(a.Data)
	p += uintptr(a.Len+len(data)-1) * Tag
	// 获取添加前数组元素最后位置的指针地址
	p2 := uintptr(a.Data)
	p2 += uintptr(a.Len-1) * Tag
	// 挪动已有在index的数据到尾巴处
	for i := index; i < a.Len+len(data); i++ {
		*(*int)(unsafe.Pointer(p - uintptr(i-index)*Tag)) = *(*int)(unsafe.Pointer(p2 - uintptr(i-index)*Tag))
	}
	//获取需要复写元素的指针地址
	p3 := uintptr(a.Data)
	p3 += uintptr(index) * Tag
	// 开始复写,安放新添加的元素
	for i, v := range data {
		*(*int)(unsafe.Pointer(p3 + uintptr(i)*Tag)) = v
	}
	a.Len += len(data)

}
// 删除数组
func (a *arr) Delect(index int) {
	a.DelectLen(index, 1)
}
// 删除全部数组
func (a *arr) DelectAll() {
	a.Len = 0
}
// 删除固定范围数组
func (a *arr) DelectLen(index, len int) {
	if a == nil || a.Data == nil {
		return
	}
	if index < 0 || len < 0 {
		return
	}
	// 获取删除起始位置元素的地址
	p := uintptr(a.Data)
	p += uintptr(index) * Tag
	// 获取删除范围后的元素地址
	p1 := p + uintptr(len)*Tag
	for i := 0; i < a.Len-len-index; i++ {
		*(*int)(unsafe.Pointer(p + uintptr(i)*Tag)) = *(*int)(unsafe.Pointer(p1 + uintptr(i)*Tag))
	}
	a.Len -= len

}
// 删除指定值
func (a *arr) DelectValue(value int) {
	for {
		i := a.ReferValue(value)
		if i == -1 {
			return
		}
		a.Delect(i)
	}
}

func NewArr() {
	arr := new(arr)
	arr.New(3, 3, 1, 2, 3)
	arr.Printf()
	arr.Add(9, 8, 7, 5, 4, 3, 88, 66, 11, 55, 001,1,2,3)
	arr.Printf()
	fmt.Println(arr.ReferIndex(5))
	fmt.Println(arr.ReferValue(7))
	arr.AddIndex(4, 12, 13, 14)
	arr.Printf()
	arr.DelectValue(1)
	arr.Printf()
	arr.DelectLen(3,4)
	arr.Printf()
	arr.Alter(1,100)
	arr.Printf()
}
