package sort

type SimpleInsertSort struct {
}

func (s SimpleInsertSort) Sort(data []interface{}, comparable Comparable) {
	dataLen := len(data)
	for i := 1; i < dataLen; i++ {
		tmp := data[i]
		j := i
		for ; j >= 1 && comparable.Compare(data[j-1], tmp) > 0; j-- {
			data[j] = data[j-1]
		}
		data[j] = tmp
	}
}
