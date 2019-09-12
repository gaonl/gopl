package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

type Car struct {
	Color byte `json:"color"`
	Price int  `json:"price"`
}

func makedemo() {

	intTyp := reflect.TypeOf(0)
	strTyp := reflect.TypeOf("")

	//反射创建map
	mapType := reflect.MapOf(strTyp, intTyp)
	v := reflect.MakeMap(mapType)
	if mapValue, ok := v.Interface().(map[string]int); ok {
		mapValue["abc"] = 100
		fmt.Println(mapValue)
	}

	//反射创建Array
	sliceType := reflect.SliceOf(intTyp)
	sliceV := reflect.MakeSlice(sliceType, 10, 10)
	if sliceValue, ok := sliceV.Interface().([]int); ok {
		for i := 0; i < 20; i++ {
			sliceValue = append(sliceValue, i)
		}

		for i, item := range sliceValue {
			fmt.Println(i, item)
		}
		fmt.Println("")
	}

	fmt.Println(reflect.ArrayOf(10, intTyp).String())             // Output: [10]int
	fmt.Println(reflect.ChanOf(reflect.BothDir, intTyp).String()) // Output: chan int
	fmt.Println(reflect.MapOf(strTyp, intTyp).String())           // Output: map[string]int
	fmt.Println(reflect.PtrTo(intTyp).String())                   // Output: *int
	fmt.Println(reflect.SliceOf(intTyp).String())                 // Output: []int
	fmt.Println(reflect.StructOf([]reflect.StructField{
		{Name: "IntA", Type: intTyp},
		{Name: "StrB", Type: strTyp},
	}).String())                                                  // Output: struct { IntA int; StrB string }
	var err error
	errTyp := reflect.TypeOf(&err).Elem()
	fmt.Println(reflect.FuncOf([]reflect.Type{intTyp}, nil, false).String())                            // Output: func(int)
	fmt.Println(reflect.FuncOf(nil, []reflect.Type{intTyp}, false).String())                            // Output: func() int
	fmt.Println(reflect.FuncOf([]reflect.Type{intTyp}, []reflect.Type{strTyp, errTyp}, false).String()) // Output: func() (string, error)
	fmt.Println(reflect.FuncOf([]reflect.Type{reflect.SliceOf(strTyp)}, nil, true).String())            // Output: func(...string)
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
	fmt.Println("1: ", reflect.ValueOf(i).Kind().String()) // Output: int
	fmt.Println("2: ", reflect.TypeOf(i).Kind().String())  // Output: int

	var pi *int
	pi = &i
	fmt.Println("3: ", reflect.ValueOf(pi).Elem().CanSet()) // Output: true   Value.Elem 和 Type.Elem()不一样
	fmt.Println("4: ", reflect.TypeOf(pi).Elem().String())  // Output: int    Type.Elem() 指的是数据指向类型

	var s string
	fmt.Println("5: ", reflect.ValueOf(s).Kind().String()) // Output: string

	var iface interface{}
	fmt.Println("6.1: ", reflect.ValueOf(iface).IsValid())     //Output: false //标记1
	fmt.Println("6: ", reflect.ValueOf(iface).Kind().String()) //Output: invalid //标记1
	//fmt.Println("7: ", reflect.TypeOf(iface).Kind().String())                 //panic.根本没有type //标记1
	fmt.Println("8: ", reflect.ValueOf(&iface).Elem().Kind().String())        //Output: interface // 标记2
	fmt.Println("9: ", reflect.ValueOf(&iface).Elem().Elem().Kind().String()) //Output: invalid // 标记3

	var ifaceNotNil interface{} = "wskj"
	fmt.Println("10: ", reflect.ValueOf(ifaceNotNil).Kind().String())                //Output: string //标记1
	fmt.Println("11: ", reflect.TypeOf(ifaceNotNil).Kind().String())                 //Output: string //标记1
	fmt.Println("12: ", reflect.ValueOf(&ifaceNotNil).Elem().Kind().String())        //Output: interface // 标记2
	fmt.Println("13: ", reflect.ValueOf(&ifaceNotNil).Elem().Elem().Kind().String()) //Output: string // 标记3

	var t time.Time
	fmt.Println(reflect.ValueOf(t).Kind().String()) //Output: struct //标记4
}

func sliceHeaderDemo(){
	slice := make([]int, 10, 100)
	fmt.Println(slice)
	sliceHeaderPtr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeaderPtr.Len = 20
	fmt.Println(slice)
}

func main() {
	//makedemo()
	//reflecttest()
	//reflectStruct()
	//sliceHeaderDemo()

	/*
	var i interface{} = "string value"
	if v,ok := i.(string); ok{
		fmt.Println(v, "is string")
	}
	i = 100
	if v,ok := i.(int); ok{
		fmt.Println(v, "is int")
	}
	 */
	var i int = 10
	var ptr = &i
	fmt.Println(unsafe.Sizeof(ptr))

}
