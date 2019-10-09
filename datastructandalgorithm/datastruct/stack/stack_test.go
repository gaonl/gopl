package stack

import (
	"fmt"
	"testing"
)

func init() {
}

func TestSimpleInsertSort(t *testing.T) {
	var s Stack = NewArrayStack()
	fmt.Println("1.", s.IsEmpty())
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	fmt.Println("2.", s.IsEmpty())
	for i := 0; i < 10; i++ {
		item, _ := s.Peek()
		fmt.Print(item)
	}
	fmt.Println()
	fmt.Println("3.", s.IsEmpty())
	for i := 0; i < 10; i++ {
		item, _ := s.Pop()
		fmt.Print(item)
	}
	fmt.Println()
	fmt.Println("4.", s.IsEmpty())
}
