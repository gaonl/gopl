package sort

type QuickSort struct {
}

func (s QuickSort) Sort(data []interface{}, comparable Comparable) {
	quickSort(data, 0, len(data)-1, comparable)
}

func quickSort(data []interface{}, start int, end int, comparable Comparable) {
	if end-start <= 10 {
		simpleInsertSort(data, start, end, comparable)
		return
	}
	item := mid(data, start, end, comparable)

	i, j := start+1, end-2
	for {
		for comparable.Compare(data[i], item) < 0 {
			i++
		}
		for comparable.Compare(data[j], item) > 0 {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
			i++
			j--
		} else {
			break
		}
	}
	data[end-1], data[i] = data[i], data[end-1]
	quickSort(data, start, i-1, comparable)
	quickSort(data, i+1, end, comparable)

}
func simpleInsertSort(data []interface{}, start int, end int, comparable Comparable) {
	for i := start + 1; i <= end; i++ {
		tmp := data[i]
		j := i
		for ; j >= start+1 && comparable.Compare(data[j-1], tmp) > 0; j-- {
			data[j] = data[j-1]
		}
		data[j] = tmp
	}
}

func mid(data []interface{}, start int, end int, comparable Comparable) interface{} {

	middle := (start + end) / 2
	if comparable.Compare(data[start], data[middle]) > 0 {
		data[start], data[middle] = data[middle], data[start]
	}
	if comparable.Compare(data[start], data[end]) > 0 {
		data[start], data[end] = data[end], data[start]
	}
	if comparable.Compare(data[middle], data[end]) > 0 {
		data[end], data[middle] = data[middle], data[end]
	}
	data[middle], data[end-1] = data[end-1], data[middle]
	return data[end-1]
}
