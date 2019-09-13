package datastructandalgorithm

import "fmt"

var arr = []int{0, 1, 2, 3, 4, 5, 6, 7}

//八皇后问题
/**
1.横向纵向都不能取一样的（横竖不是一条直线）
	a.横向0-7取排列组合，用到排列组合算法，比如[4,2,5,7,6,8,1,3]是40320种可能的其中一种
	b.纵向去固定排列(不一定要有序，程序实就用有序的吧)[0,1,2,3,4,5,6,7,]
2.去位置相同的倆数，相加后都不能一样，相减后也都不能一样（两个斜线不是一条直线）
	[4,2,5,7,6,8,1,3]
	[0,1,2,3,4,5,6,7]
	比如：4+0=4 2+1=3 5+2=7 7+3=10 6+4=10 8+5=13 1+6=7 3+7=10  有三个10，两个7 不符合
	比如：4-0=4 2-1=1 5-2=3 7-3=4  6-4=2  8-5=3  1-6=-5 3-7=-4  有两个4，两个3 不符合
	要两个都符合，这个组合才是八皇后的一个解
*/
func EightQueues(arr []int) {
	for PermuteOnce(arr) {
		sumMap := make(map[int]bool)
		subMap := make(map[int]bool)
		plate := [8][8]byte{}
		for i, item := range arr {
			plate[i][item] = 1

			sum := i + item
			sub := i - item
			if sumMap[sum] {
				break
			}
			if subMap[sub] {
				break
			}
			sumMap[sum] = true
			subMap[sub] = true
			if i == 7 {
				for _, row := range plate {
					fmt.Println(row)
				}
				fmt.Println("-----------")
				plate = [8][8]byte{}
			}
		}
	}
}