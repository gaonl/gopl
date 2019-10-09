package stack

type Stack interface {
	Pop() (interface{}, bool)
	Push(interface{}) bool
	Peek() (interface{}, bool)
	Empty() bool
	IsEmpty() bool
}
