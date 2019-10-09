package stack

const initCap = 100

type arrayStack struct {
	len  int
	cap  int
	data [initCap]interface{}
}

func NewArrayStack() *arrayStack {
	return &arrayStack{
		len: 0,
		cap: initCap,
	}
}

func (a *arrayStack) Pop() (interface{}, bool) {
	if a.len == 0 {
		return nil, false
	}
	item := a.data[a.len-1]
	a.len--
	return item, true
}
func (a *arrayStack) Peek() (interface{}, bool) {
	if a.len == 0 {
		return nil, false
	}
	return a.data[a.len-1], true
}

func (a *arrayStack) Push(item interface{}) bool {
	if a.len == a.cap {
		return false
	}
	a.data[a.len] = item
	a.len++
	return true
}

func (a *arrayStack) Empty() bool {
	a.len = 0
	return true
}

func (a *arrayStack) IsEmpty() bool {
	return a.len == 0
}
