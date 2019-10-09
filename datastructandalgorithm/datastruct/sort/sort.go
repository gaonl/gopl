package sort

const (
	SIMPLE_INSERT_SORT string = "SIMPLE_INSERT_SORT"
	SHELL_SORT         string = "SHELL_SORT"
	BUBBLE_SORT        string = "BUBBLE_SORT"
	HEAP_SORT          string = "HEAP_SORT"
	MERGE_SORT         string = "MERGE_SORT"
	QUICK_SORT         string = "QUICK_SORT"
)

type Sorter interface {
	Sort([]interface{}, Comparable)
}

type SorterManager struct {
	sorterMap map[string]Sorter
}

func NewSorterManager() SorterManager {
	return SorterManager{sorterMap: make(map[string]Sorter, 10)}
}

func (s SorterManager) Register(key string, sorter Sorter) {
	s.sorterMap[key] = sorter
}

func (s SorterManager) Sort(key string, data []interface{}, comparable Comparable) {
	sorter := s.sorterMap[key]
	sorter.Sort(data, comparable)
}
