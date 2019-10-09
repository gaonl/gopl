package sort

type BubbleSort struct {
}

func (s BubbleSort) Sort(data []interface{}, comparable Comparable) {
	dataLen := len(data)
	if dataLen < 2 {
		return
	}
	for j := dataLen - 1; j >= 0; j-- {
		flag := false
		for i := 0; i < j; i++ {
			if comparable.Compare(data[i], data[i+1]) > 0 {
				data[i], data[i+1] = data[i+1], data[i]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}
