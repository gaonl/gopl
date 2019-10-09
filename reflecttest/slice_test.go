package main

import (
	"fmt"
	"testing"
)

var s = make([]int, 0, 10)

func TestStruct(t *testing.T) {

	fmt.Printf("---->%p\n", &s)
	for i := 0; i < 20; i++ {
		fmt.Printf("%p\n", &s)
		_ = append(s, i)
		fmt.Printf("%p\n", &s)
	}
	fmt.Println(s)
}
