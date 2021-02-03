package stl

func heapSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	//堆化
	buildHeap(nums, l)

	for i := l - 1; i >= 0; i-- {
		//大顶堆后，最大的在头部，放到尾部去
		nums[i], nums[0] = nums[0], nums[i]
		//对前面的重新堆化
		l--
		heapDown(nums, 0, l)
	}
	return nums
}

func heapDown(nums []int, i, arrLen int) {
	left, right := i*2+1, i*2+2
	max := i
	if left < arrLen && nums[max] < nums[left] {
		max = left
	}
	if right < arrLen && nums[max] < nums[right] {
		max = right
	}
	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		heapDown(nums, max, arrLen)
	}
}

func buildHeap(nums []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapDown(nums, i, arrLen)
	}
}
