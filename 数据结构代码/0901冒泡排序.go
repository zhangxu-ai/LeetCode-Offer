package main

import "fmt"

//冒泡法初级，将第i个元素的值与其后的所有的len-i个值进行比较
func BubbleSort0(sl []int) {
	for i := 0; i < len(sl); i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i] > sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
}

//冒泡法正宗版,从序列末尾22比较，一直将最小值冒泡到最顶端
func BubbleSort1(sl []int) {
	for i := 0; i < len(sl); i++ {
		for j := len(sl) - 1; j > i; j-- {
			if sl[j] < sl[j-1] {
				sl[j-1], sl[j] = sl[j], sl[j-1]
			}
		}
	}
}

//冒泡法优化，在i变化到某一步时，若已经达到了平衡，
//则 swap函数不被调用，isSorted 也不会被置为false，从而退出最外层的循环
func BubbleSort2(sl []int) {
	//默认没有被排序好
	isSorted := false
	//只有没有排序的时候才会进入循环
	for i := 0; i < len(sl) && !isSorted; i++ {
		//假设已近排好序了
		isSorted = true
		for j := len(sl) - 1; j > i; j-- {
			if sl[j] < sl[j-1] {
				sl[j], sl[j-1] = sl[j-1], sl[j]

				//进行了交换操作，说明还有没排序好的
				isSorted = false
			}
		}
	}
}
func main() {
	o := []int{2, 1, 3, 5, 6, 1, 8, 10, 7}
	//BubbleSort0(o)
	BubbleSort2(o)
	fmt.Println(o)
}
