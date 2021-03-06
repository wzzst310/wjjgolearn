package heap

import (
	"fmt"
	common "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_common"
	"testing"
)

func TestHeap(t *testing.T) {
	// test1()
	test2()
}
func test2() {
	heap := NewBinaryHeap(common.IntRevertComparator)
	heap.Add(9)
	heap.Add(79)
	heap.Add(2)
	heap.Add(16)
	heap.Add(27)
	heap.Add(33)
	heap.Add(6)
	for !heap.IsEmpty() {
		fmt.Printf("%#v", heap)
		fmt.Println(heap.Remove())
	}
	//heap.Remove()
	//heap.Remove()
	//heap.Replace(1)
	//fmt.Printf("%#v", heap)
}

func test1() {
	// 最小堆
	heap := NewBinaryHeap(common.IntComparator)
	// 找出前top个元素
	top := 5
	elements := []int{68, 83, 97, 89, 50, 62, 9, 15, 44, 38, 57, 98, 21, 64, 85, 63, 87, 13, 42, 90, 4, 40, 2, 36, 17, 48, 5, 19, 99, 12}
	for i := 0; i < len(elements); i++ {
		if i < top {
			heap.Add(elements[i])
		} else if elements[i]-heap.Get().(int) > 0 {
			heap.Replace(elements[i])
		}
	}
	fmt.Printf("%#v", heap)
}
