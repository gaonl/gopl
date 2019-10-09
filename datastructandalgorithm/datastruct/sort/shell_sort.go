package sort

type ShellSort struct {
}

func (s ShellSort) Sort(data []interface{}, comparable Comparable) {
	dataLen := len(data)
	//希尔排序中的增量，简单的用这个
	//好的增量书说是这个，分析很复杂，就不深究原理了[1,5,19,41,109,....]
	//公式是：9*4^i -9*2^i +1 或者 4^i -3*2^i +1
	for incr := dataLen / 2; incr >= 1; incr /= 2 {
		for i := 0; i < incr; i++ {
			insertPart(data, i, incr, comparable)
		}
	}
}

//shell排序中，每一趟就是一个插入排序
func insertPart(data []interface{}, start int, incr int, comparable Comparable) {
	dataLen := len(data)
	for i := incr + start; i < dataLen; i += incr {
		tmp := data[i]
		j := i
		for ; j >= incr + start && comparable.Compare(data[j-incr], tmp) > 0; j-=incr {
			data[j] = data[j-incr]
		}
		data[j] = tmp
	}
}
