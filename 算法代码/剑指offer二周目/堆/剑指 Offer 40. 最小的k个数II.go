package main

import "fmt"

func getLeastNumbersII(arr []int, k int) []int {
	if len(arr) == 0 || k >= len(arr) {
		return arr
	}
	if k == 0 {
		return nil
	}
	heap := arr[:k]
	fmt.Println(heap)
	makeHeap(heap)
	fmt.Println(heap)

	for i := k; i < len(arr); i++ {
		if arr[i] < heap[0] {
			heap[0] = arr[i]
			heapify(heap, 0)
		}
	}
	return heap
}

func makeHeap(arr []int) {
	l := len(arr)
	lastNode := l - 1
	lastParent := (lastNode - 1) / 2
	for i := lastParent; i >= 0; i-- {
		heapify(arr, i)
	}
}

func heapify(arr []int, i int) {
	left := i*2 + 1
	right := left + 1
	max := i
	if left < len(arr) && arr[left] > arr[max] {
		max = left
	}
	if right < len(arr) && arr[right] > arr[max] {
		max = right
	}

	if max != i {
		arr[i], arr[max] = arr[max], arr[i]
		heapify(arr, max)
	}
}
func main() {
	a := []int{0, 1, 2, 1, 3, 0, 1, -1}
	fmt.Println(getLeastNumbersII(a, 3))
}
