package stl

func InsertSort(sl []int) {
	for i := 1; i < len(sl); i++ {
		//v是当前元素
		v := sl[i]
		//index指向前一个元素
		index := i - 1
		//将当前元素一直和其前面的元素进行比较，当前面的比当前元素大时
		//将前面的元素向后移动一位，直到
		//1. 比较到第一个元素，此时index=-1，
		//2. index元素大于v，且index+1的元素小于v，说明v在index之后，
		for index >= 0 && sl[index] > v {
			sl[index+1] = sl[index]
			index--
		}
		//而index+1的元素已经复制一份到index+2上了（后移了），所以可以将v放到index+1上了
		sl[index+1] = v
	}
}
