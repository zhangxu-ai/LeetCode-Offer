package main

import "fmt"

func SelectSort(sl []int) {
	var min = 0
	for i := 0; i < len(sl); i++ {
		min = i
		//不断寻找最小值
		for j := i + 1; j < len(sl); j++ {
			if sl[j] < sl[min] {
				min = j
			}
		}
		//将最小值放到最前面
		if min != i {
			sl[min], sl[i] = sl[i], sl[min]
		}
	}
}
func main() {
	o := []int{2, 1, 3, 5, 6, 1, 8, 10, 7}
	SelectSort(o)
	fmt.Println(o)
}
