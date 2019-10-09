package sort

type HeapSort struct {
}

func (s HeapSort) Sort(data []interface{}, comparable Comparable) {
	buildHeap(data, comparable)

	n := len(data) - 1

	for i := n; i >= 1; i-- {
		max2end(data, i, comparable)
	}

}

func max2end(heap []interface{}, n int, comparable Comparable) {
	heap[1], heap[n] = heap[n], heap[1]
	down1(heap, 1, n-1, comparable)
}

func buildHeap(data []interface{}, comparable Comparable) {
	n := len(data) - 1
	for i := n / 2; i >= 1; i-- {
		down1(data, i, n, comparable)
	}
}

//构建当个节点的堆  data第0个位置不用，i要构建节点的索引位置，n当前堆元素个数
func down1(data []interface{}, i int, n int, comparable Comparable) {
	leftIdx := i * 2
	for leftIdx <= n {
		rightIdx := leftIdx + 1
		child := leftIdx

		if rightIdx <= n && comparable.Compare(data[rightIdx], data[leftIdx]) > 0 {
			child = rightIdx
		}

		if comparable.Compare(data[child], data[i]) > 0 {
			data[child], data[i] = data[i], data[child]
			leftIdx = child * 2
			i = child
		} else {
			break
		}
	}
}
