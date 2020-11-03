package main

func (sl *SqList) SelectSort() {
	var min = 0
	for i := 0; i < sl.length; i++ {
		min = i
		//不断寻找最小值
		for j := i + 1; j < sl.length; j++ {
			if sl.data[j] < sl.data[min] {
				min = j
			}
		}
		//将最小值放到最前面
		if min != i {
			sl.Swap(min, i)
		}
	}
}
