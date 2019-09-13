package main

import (
	"bytes"
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

type Car struct {
	Color byte `json:"color"`
	Price int  `json:"price"`
}

func reflectStruct() {
	var c Car
	t := reflect.TypeOf(c)
	field, ok := t.FieldByName("Color")
	if ok {
		fmt.Println(field.Name)
		fmt.Println(field.Tag)
	}

	v := reflect.ValueOf(c)
	value := v.FieldByName("Colors")
	fmt.Println(value.IsValid())
}

func reflecttest() {
	var i int
	println("1: ", reflect.ValueOf(i).Kind().String()) // Output: int
	println("2: ", reflect.TypeOf(i).Kind().String())  // Output: int

	var pi *int
	pi = &i
	println("3: ", reflect.ValueOf(pi).Elem().CanSet()) // Output: true   Value.Elem 和 Type.Elem()不一样
	println("4: ", reflect.TypeOf(pi).Elem().String())  // Output: int    Type.Elem() 指的是数据指向类型

	var s string
	println("5: ", reflect.ValueOf(s).Kind().String()) // Output: string

	var iface interface{}
	println("6.1: ", reflect.ValueOf(iface).IsValid())     //Output: false //标记1
	println("6: ", reflect.ValueOf(iface).Kind().String()) //Output: invalid //标记1
	//println("7: ", reflect.TypeOf(iface).Kind().String())                 //panic.根本没有type //标记1
	println("8: ", reflect.ValueOf(&iface).Elem().Kind().String())        //Output: interface // 标记2
	println("9: ", reflect.ValueOf(&iface).Elem().Elem().Kind().String()) //Output: invalid // 标记3

	var ifaceNotNil interface{} = "wskj"
	println("10: ", reflect.ValueOf(ifaceNotNil).Kind().String())                //Output: string //标记1
	println("11: ", reflect.TypeOf(ifaceNotNil).Kind().String())                 //Output: string //标记1
	println("12: ", reflect.ValueOf(&ifaceNotNil).Elem().Kind().String())        //Output: interface // 标记2
	println("13: ", reflect.ValueOf(&ifaceNotNil).Elem().Elem().Kind().String()) //Output: string // 标记3

	var t time.Time
	println(reflect.ValueOf(t).Kind().String()) //Output: struct //标记4
}

func main() {
	//reflecttest()
	//reflectStruct()

	/*
		var p *int
		fmt.Println(reflect.TypeOf(p))                   //*int
		fmt.Println(reflect.ValueOf(p))                  //<nil>
		fmt.Println(reflect.ValueOf(p).Elem())           //<invalid reflect.Value>
		fmt.Println(reflect.ValueOf(p).Elem().IsValid()) //false
	*/

	slicestr := []byte{'a', 'b', 'c', 'd', 'e'}
	str := (*string)(unsafe.Pointer(&slicestr))
	fmt.Println(*str) //abcde
	slicestr[2] = '*'
	fmt.Println(*str) //ab*de

	buf := bytes.Buffer{}
	buf.Bytes()
}
