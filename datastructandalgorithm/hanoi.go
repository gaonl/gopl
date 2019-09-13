package datastructandalgorithm

import "fmt"

//从a 借助b  移动到c
func Move(n int, a, b, c byte) {
	if n == 1 {
		fmt.Printf("%c to %c\n", a, c)
		return
	}
	Move(n-1, a, c, b)
	fmt.Printf("%c to %c\n", a, c)
	Move(n-1, b, a, c)
}

func main() {
	Move(5, 'a', 'b', 'c')
}
