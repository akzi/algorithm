package pairingheap

import (
	"sort"
	"testing"
)

type (
	Int int
)

func (first Int) Less(second Item) bool {
	return first < second.(Int)
}

func TestHeap_Range(t *testing.T) {
	heap := New()
	ints := []int{9, 10, 11, 12, 8, 13, 14, 7, 15, 16}
	for _, v := range ints {
		heap.Insert(Int(v))
	}
	sort.Ints(ints)
	heap.Range(func(item Item) bool {
		i := sort.Search(len(ints), func(i int) bool {
			return  ints[i] >= int(item.(Int))
		})
		if i == len(ints) {
			t.Fatal("not find",int(item.(Int)),i,ints)
		}
		return true
	})
}

func TestDeleteMin(t *testing.T) {
	heap := New()
	ints := []int{9, 10, 11, 12, 8, 13, 14, 7, 15, 16}
	for _, v := range ints {
		heap.Insert(Int(v))
	}

	sort.Ints(ints)
	for len(ints) != 0 {
		val := heap.FindMin().(Int)
		if int(val) != ints[0] {
			t.Fatal("error", int(val), ints[0])
		}
		ints = ints[1:]
		heap.DeleteMin()
	}
}
