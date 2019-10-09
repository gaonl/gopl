package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

var sorterManager SorterManager

func init() {
	sorterManager = NewSorterManager()
	sorterManager.Register(SIMPLE_INSERT_SORT, SimpleInsertSort{})
	sorterManager.Register(SHELL_SORT, ShellSort{})
	sorterManager.Register(BUBBLE_SORT, BubbleSort{})
	sorterManager.Register(HEAP_SORT, HeapSort{})
	sorterManager.Register(QUICK_SORT, QuickSort{})
}

func TestSimpleInsertSort(t *testing.T) {
	intData := makeRandomIntSlice(10000, 10000)
	fmt.Println(intData)
	fmt.Println("sorted(before):", testSorted(intData, IntComparable{}))

	sorterManager.Sort(SIMPLE_INSERT_SORT, intData, IntComparable{})
	fmt.Println(intData)
	fmt.Println("sorted(after):", testSorted(intData, IntComparable{}))

}
func TestShellSort(t *testing.T) {
	intData := makeRandomIntSlice(10000, 10000)
	fmt.Println(intData)
	fmt.Println("sorted(before):", testSorted(intData, IntComparable{}))

	sorterManager.Sort(SHELL_SORT, intData, IntComparable{})
	fmt.Println(intData)
	fmt.Println("sorted(after):", testSorted(intData, IntComparable{}))
}
func TestBubbleSort(t *testing.T) {
	intData := makeRandomIntSlice(10000, 10000)
	fmt.Println(intData)
	fmt.Println("sorted(before):", testSorted(intData, IntComparable{}))

	sorterManager.Sort(BUBBLE_SORT, intData, IntComparable{})
	fmt.Println("bubble sort", intData)
	fmt.Println("sorted(after):", testSorted(intData, IntComparable{}))
}
func TestHeapSort(t *testing.T) {
	intData := makeRandomIntSlice(10000, 10000)
	intData[0] = -1
	fmt.Println(intData)
	fmt.Println("sorted(before):", testSorted(intData, IntComparable{}))

	sorterManager.Sort(HEAP_SORT, intData, IntComparable{})
	fmt.Println("heap sort", intData)
	fmt.Println("sorted(after):", testSorted(intData, IntComparable{}))
}

func TestQuickSort(t *testing.T) {
	intData := makeRandomIntSlice(10000, 10000)
	fmt.Println(intData)
	fmt.Println("sorted(before):", testSorted(intData, IntComparable{}))

	sorterManager.Sort(QUICK_SORT, intData, IntComparable{})
	fmt.Println(intData)
	fmt.Println("sorted(after):", testSorted(intData, IntComparable{}))
}

func makeRandomIntSlice(n, max int) []interface{} {
	var intData = make([]interface{}, n, n)
	for i := 0; i < n; i++ {
		intData[i] = rand.Intn(max)
	}
	return intData
}

func testSorted(sorted []interface{}, comparable Comparable) bool {
	max := sorted[0]
	for i := 1; i < len(sorted); i++ {
		if comparable.Compare(sorted[i], max) < 0 {
			return false
		}
		max = sorted[i]
	}
	return true;
}
