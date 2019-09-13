package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string
	var s []int
	var m map[string]int
	var f func(int, int) int
	var i interface{}
	var c chan int
	var ptr *int
	fmt.Println(str == "")
	fmt.Println(s == nil)
	fmt.Println(m == nil)
	fmt.Println(f == nil)
	fmt.Println(i == nil)
	fmt.Println(c == nil)
	fmt.Println(ptr == nil)

	fmt.Println("--------------")
	fmt.Println(reflect.TypeOf(str)) //2*byte 即输出16  指向底层数组指针（Data）+len(长度)，可以看reflect.StringHeader
	fmt.Println(reflect.TypeOf(s))   //3*byte 即输出24  指向底层数组指针+len+cap，可以看reflect.SliceHeader
	fmt.Println(reflect.TypeOf(m))   //1*byte 即输出8 本质也是指针
	fmt.Println(reflect.TypeOf(f))   //1*byte 即输出8 本质也是指针
	fmt.Println(reflect.TypeOf(i))   //2*byte 即输出8 一个指向类型（Type）一个指向值（Value）反射用到的就是这个
	fmt.Println(reflect.TypeOf(c))   //1*byte 即输出8 本质也是指针
	fmt.Println(reflect.TypeOf(ptr)) //1*byte 即输出8 指针

	fmt.Println("--------------")
	fmt.Println(unsafe.Sizeof(str))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(m))
	fmt.Println(unsafe.Sizeof(f))
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(ptr))
}
