package main

//冒泡法初级，将第i个元素的值与其后的所有的len-i个值进行比较
func (sl *SqList) BubbleSort0() {
	for i := 0; i < sl.length; i++ {
		for j := i + 1; j < sl.length; j++ {
			if sl.data[i] > sl.data[j] {
				sl.Swap(sl.data[i], sl.data[j])
			}
		}
	}
}

//冒泡法正宗版,从序列末尾22比较，一直将最小值冒泡到最顶端
func (sl *SqList) BubbleSort1() {
	for i := 0; i < sl.length; i++ {
		for j := sl.length - 1; j > i; j-- {
			if sl.data[j] < sl.data[j-1] {
				sl.Swap(sl.data[j], sl.data[j-1])
			}
		}
	}
}

//冒泡法优化，在i变化到某一步时，若已经达到了平衡，
//则 swap函数不被调用，isSorted 也不会被置为false，从而退出最外层的循环
func (sl *SqList) BubbleSort2() {
	//默认没有被排序好
	isSorted := false
	//只有没有排序的时候才会进入循环
	for i := 0; i < sl.length && !isSorted; i++ {
		//假设已近排好序了
		isSorted = true
		for j := sl.length - 1; j > i; j-- {
			if sl.data[j] < sl.data[j-1] {
				sl.Swap(sl.data[j], sl.data[j-1])
				//进行了交换操作，说明还有没排序好的
				isSorted = false
			}
		}
	}
}
