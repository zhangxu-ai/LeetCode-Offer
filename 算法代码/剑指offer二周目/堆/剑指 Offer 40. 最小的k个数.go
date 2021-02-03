package main

import "fmt"

//快速排序
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k >= len(arr) {
		return arr
	}
	if k == 0 {
		return nil
	}
	quickSort(arr)
	fmt.Println(arr)
	return arr[:k]
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	p := arr[0]
	low, big := 0, len(arr)-1
	for low < big {
		for low < big && arr[big] >= p {
			big--
		}
		arr[low] = arr[big]
		for low < big && arr[low] <= p {
			low++
		}
		arr[big] = arr[low]
	}
	arr[low] = p
	quickSort(arr[:low])
	quickSort(arr[low+1:])
}

func main() {
	a := []int{0, 1, 2, 1, 3, 0, 1, -1}
	fmt.Println(getLeastNumbers(a, 3))
}
