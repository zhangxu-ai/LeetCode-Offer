package stl

func LSDSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	max := nums[0]
	maxDigit := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	for max != 0 {
		max /= 10
		maxDigit++
	}
	buks := make([][]int, 10)
	for i := 0; i < 10; i++ {
		buks[i] = []int{}
	}
	for i, c := 1, 1; i <= maxDigit; i++ {
		for j := 0; j < len(nums); j++ {
			b := (nums[j] / c) % 10
			buks[b] = append(buks[b], nums[j])
		}
		c *= 10
		k := 0
		for i := 0; i <= 9; i++ {
			for j := 0; j < len(buks[i]); j++ {
				nums[k] = buks[i][j]
				k++
			}
			buks[i] = []int{}
		}
	}
	return nums
}
