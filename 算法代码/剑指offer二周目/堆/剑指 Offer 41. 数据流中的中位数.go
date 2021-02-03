package main

import "fmt"

func heapMaxDown(arr []int, id int) {
	left := id*2 + 1
	right := left + 1
	max := id
	if left < len(arr) && arr[left] > arr[max] {
		max = left
	}
	if right < len(arr) && arr[right] > arr[max] {
		max = right
	}
	if max != id {
		arr[max], arr[id] = arr[id], arr[max]
		heapMaxDown(arr, max)
	}
}

func heapMinDown(arr []int, id int) {
	left := id*2 + 1
	right := left + 1
	min := id
	if left < len(arr) && arr[left] < arr[min] {
		min = left
	}
	if right < len(arr) && arr[right] < arr[min] {
		min = right
	}
	if min != id {
		arr[min], arr[id] = arr[id], arr[min]
		heapMinDown(arr, min)
	}
}
func buildHeap(arr []int, isMax bool) {
	l := len(arr)
	lastNode := l - 1
	lastParent := (lastNode - 1) / 2
	if isMax {
		for i := lastParent; i >= 0; i-- {
			heapMaxDown(arr, i)
		}
	} else {
		for i := lastParent; i >= 0; i-- {
			heapMinDown(arr, i)
		}

	}
}

type MedianFinder struct {
	maxHeap []int
	minHeap []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: []int{},
		minHeap: []int{},
	}
}

func (f *MedianFinder) AddNum(num int) {
	if len(f.maxHeap) == 0 {
		f.maxHeap = append(f.maxHeap, num)
		return
	}

	if len(f.maxHeap) == len(f.minHeap) {
		f.minHeap = append(f.minHeap, num)
		buildHeap(f.minHeap, false)
		f.maxHeap = append(f.maxHeap, f.minHeap[0])
		buildHeap(f.maxHeap, true)
		f.minHeap = f.minHeap[1:]
		buildHeap(f.minHeap, false)
	} else {
		f.maxHeap = append(f.maxHeap, num)
		buildHeap(f.maxHeap, true)
		f.minHeap = append(f.minHeap, f.maxHeap[0])
		buildHeap(f.minHeap, false)
		f.maxHeap = f.maxHeap[1:]
		buildHeap(f.maxHeap, true)

	}
}

func (f *MedianFinder) FindMedian() float64 {
	if len(f.maxHeap) == 0 {
		return 0.0
	}
	if len(f.maxHeap) == len(f.minHeap) {
		return (float64(f.maxHeap[0]) + float64(f.minHeap[0])) / 2
	} else {
		return float64(f.maxHeap[0])
	}
}

func main() {
	f := Constructor()
	a := []int{6, 10, 2, 6, 5}
	for _, i := range a {
		f.AddNum(i)
		fmt.Println(f.maxHeap, f.minHeap, f.FindMedian())
	}
}
