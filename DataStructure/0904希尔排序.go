package stl

func ShellSort(sl []int) {
	//确定初始间隔为长度/2,之后每一轮减半
	//当inc=1时，就是进行了一轮直接插入排序//直接插入排序可以看做间隔为1的希尔排序
	for inc := len(sl) / 2; inc >= 1; inc /= 2 {
		//每一轮，将i处的元素与其前面间隔为inc的元素们看做一个子序列,并对其进行快速排序
		for i := inc; i < len(sl); i++ {
			j := i
			tmp := sl[i]
			//将当前元素一直和其前面间隔为inc的元素进行比较，当前面的比当前元素大时
			//将前面的元素向后移动inc位
			for ; j >= inc && tmp < sl[j-inc]; j -= inc {
				sl[j] = sl[j-inc]
			}
			//直到
			//1. j的前inc个元素超过了序列的首部，此时j-inc<0
			//2. j-inc元素大于tmp，且j-inc-inc的元素小于tmp，说明tmp在j-inc之后，
			//然后将tmp插入到当前j的位置上
			sl[j] = tmp
		}
	}
}
