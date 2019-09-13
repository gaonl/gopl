package datastructandalgorithm

//求解next数组
func next(sub string) []int {
	subSlice := []byte(sub)
	subLength := len(subSlice)
	next := make([]int, subLength)
	next[0] = 0

	/*
		求解相同前缀和后移获取next数组可以同时跑
		这里为了直观，分两步（优化见moreEffectiveNext函数）
	*/
	//1.求解相同前缀、后缀长度数组
	for i, j := 0, 1; j < subLength; j++ {
		if subSlice[i] == subSlice[j] {
			next[j] = i + 1
			i++
		} else {
			i = 0
		}
	}

	//2.后移，第一位补-1
	for i := subLength - 1; i > 0; i-- {
		next[i] = next[i-1]
	}
	next[0] = -1
	return next
}

//求解next数组(优化的next数组)
func moreEffectiveNext(sub string) []int {
	subSlice := []byte(sub)
	subLength := len(subSlice)
	next := make([]int, subLength)
	next[0] = -1
	next[1] = 0
	for i, j := 0, 1; j < subLength-1; j++ {
		if subSlice[i] == subSlice[j] {
			next[j+1] = i + 1
			i++
		} else {
			i = 0
		}
	}
	return next
}

func SubString(str, sub string) (n int, ok bool) {
	next := moreEffectiveNext(sub)
	strSlice := []byte(str)
	subSlice := []byte(sub)
	strLength := len(strSlice)
	subLength := len(subSlice)

	i, j := 0, 0
	for i < strLength && j < subLength {
		if j == -1 || strSlice[i] == subSlice[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == subLength {
		return i - j, true
	}

	return -1, false
}

//字符串匹配、朴素算法
func SubStringNormal(str, sub string) (n int, ok bool) {
	strSlice := []byte(str)
	subSlice := []byte(sub)
	strLength := len(strSlice)
	subLength := len(subSlice)
	i, j := 0, 0
	nexti := i + 1
	for i < strLength && j < subLength {
		if strSlice[i] == subSlice[j] {
			i++
			j++
		} else {
			j = 0
			i = nexti
			nexti++
		}
	}
	if j == subLength {
		return i - j, true
	}
	return -1, false
}