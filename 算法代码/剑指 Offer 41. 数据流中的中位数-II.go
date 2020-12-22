package main

import (
	"container/heap"
	"fmt"
)

//使用大小堆
//53，33

type (
	MinHeap []int
	MaxHeap []int
)

//------小根堆
func (m MinHeap) Len() int {
	return len(m)
}
func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *MinHeap) Pop() interface{} {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return v
}

//------大根堆
func (m MaxHeap) Len() int {
	return len(m)
}
func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *MaxHeap) Pop() interface{} {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return v
}

type MedianFinderH struct {
	min *MinHeap
	max *MaxHeap
}

func ConstructorMH() MedianFinderH {
	min := new(MinHeap)
	max := new(MaxHeap)
	heap.Init(min)
	heap.Init(max)
	return MedianFinderH{
		min: min,
		max: max,
	}
}
func (h MedianFinderH) AddNum(num int) {
	// 压入数据时有两种情况；
	// （1）原有数据为偶数个，则有 leftheap.len() == rightheap.len()；
	//     此时将数据放入leftheap，需先将数据放在rightheap，再从rightheap中pop出一个元素，将其放入leftheap；
	//     上面做法的高明之处 在于 省去了判断 num 与 rightheap最小值 及 leftheap最大值 的比较，
	//    通过比较判断num应该放在哪个heap，
	//     然后再根据左右heap的长度，对左右heap进行调整
	//  (2) 原有数据为奇数个，则有 leftheap.len() == rightheap.len() + 1
	// (通过step 1 来保证当为奇数时，left总比right 大 1，这样可以简化判断)
	//     类似step 1 需先将数据放在leftheap，再从leftheap中pop出一个元素，将其放入rightheap中
	if h.min.Len() == h.max.Len() {
		heap.Push(h.min, num)
		heap.Push(h.max, heap.Pop(h.min))
	} else {
		heap.Push(h.max, num)
		heap.Push(h.min, heap.Pop(h.max))
	}
}

func (h MedianFinderH) findMedian() float64 {
	if h.min.Len() == h.max.Len() {
		return float64((*(h.min))[0]+(*(h.max))[0]) / 2
	} else {
		return float64((*(h.max))[0])
	}
}
func main() {
	mf := ConstructorMH()
	mf.AddNum(6)
	mf.AddNum(10)
	fmt.Println(mf.findMedian())
}
