package sort

import (
	"fmt"
	"strings"
)

type Comparable interface {
	Compare(interface{}, interface{}) int
}
type IntComparable struct {
}

func (i IntComparable) Compare(a interface{}, b interface{}) int {
	ia, ok := a.(int)
	if !ok {
		panic(fmt.Sprintf("data %v is not a int.", a))
	}
	ib, ok := b.(int)
	if !ok {
		panic(fmt.Sprintf("data %v is not a int.", b))
	}

	return ia - ib
}

type StringComparable struct {
}

func (i StringComparable) Compare(a interface{}, b interface{}) int {
	sa, ok := a.(string)
	if !ok {
		panic(fmt.Sprintf("data %v is not a string.", a))
	}
	sb, ok := b.(string)
	if !ok {
		panic(fmt.Sprintf("data %v is not a string.", b))
	}

	return strings.Compare(sa, sb)
}
