package datastructandalgorithm

import "fmt"

func swap(i, j *int) {
	*i, *j = *j, *i
}

//排列组合(字典法)
/*
生成给定全排列的下一个排列 所谓一个的下一个就是这一个与下一个之间没有字典顺序中相邻的字符串。
这就要求这一个与下一个有尽可能长的共同前缀，也即变化限制在尽可能短的后缀上。

例子:839647521的下一个排列.
从最右开始,找到第一个比右边小的数字4(因为4<7，而7>5>2>1),
再从最右开始,找到4右边比4大的数字5(因为4>2>1而4<5),交换4、5,【此时5右边为7421,倒置为1247【不要忘了】】
,即得下一个排列:839651247.用此方法写出全排列的非递归算法如下
*/
func Permute1(sorted []int) {
	fmt.Println(sorted)
	for PermuteOnce(sorted) {
		fmt.Println(sorted)
	}
}
func PermuteOnce(sorted []int) bool {
	maxIndex := len(sorted) - 1
	//for {
	i, j := maxIndex-1, maxIndex
	//从最右开始,找到第一个比右边小的数字
	for i >= 0 {
		if sorted[i] < sorted[i+1] {
			break
		}
		i--
	}
	//已经找不到【第一个比右边小的数字】证明已经搜完了
	if i < 0 {
		return false
	}

	//再从最右开始,找到i右边比i大的数字，如果有就交换
	for j > i {
		if sorted[j] > sorted[i] {
			sorted[j], sorted[i] = sorted[i], sorted[j]
			//倒置i(不包含i)后边的数组
			start := i + 1
			end := maxIndex
			for start != end && start < end {
				sorted[start], sorted[end] = sorted[end], sorted[start]
				start++
				end--
			}
			break
		}
		j--
	}
	//}
	return true
}

//排列组合(递归法)
/**
例子：[0,1,2,3,4,5,6,7]
过程：
[0,[1,2,3,4,5,6,7排列组合]]
[1,[0,2,3,4,5,6,7排列组合]]
[2,[1,0,3,4,5,6,7排列组合]]
[3,[1,2,0,4,5,6,7排列组合]]
[4,[1,2,3,0,5,6,7排列组合]]
[5,[1,2,3,4,0,6,7排列组合]]
[6,[1,2,3,4,5,0,7排列组合]]
[7,[1,2,3,4,5,6,0排列组合]]
*/
func Permute(arr []int, start, end int) {
	if start == end {
		fmt.Println(arr)
	} else {
		for i := start; i <= end; i++ {
			swap(&(arr[i]), &(arr[start]))
			Permute(arr, start+1, end)
			swap(&(arr[i]), &(arr[start]))
		}
	}
}
