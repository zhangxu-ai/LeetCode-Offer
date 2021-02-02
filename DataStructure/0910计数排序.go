package stl

func countingSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	//计算最大值
	max := nums[0]
	for i := 1; i < l; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	tmp := make([]int, max+1)

	for i := 0; i < l; i++ {
		tmp[nums[i]] += 1
	}
	index := 0
	for i := 0; i < max+1; i++ {
		for tmp[i] > 0 {
			nums[index] = i
			index++
			tmp[i] -= 1
		}
	}
	return nums
}
