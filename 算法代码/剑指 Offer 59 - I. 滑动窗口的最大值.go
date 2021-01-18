package main

//算是优化后的暴力法吧
var ns []int

func maxSlidingWindow(nums []int, k int) []int {
	ns = nums
	res := make([]int, 0, len(nums)-k+1)
	i, j, maxIndex := 0, k-1, 0
	for j <= len(nums)-1 {
		if i == 0 {
			maxIndex = max(i, k)
		} else {
			if nums[j] >= nums[maxIndex] {
				maxIndex = j
			} else {
				if maxIndex == i-1 {
					maxIndex = max(i, k)
				}
			}
		}
		res = append(res, nums[maxIndex])
		i++
		j++
	}
	return res
}

func max(start, len int) int {
	maxIndex := start
	for i := start; i < start+len; i++ {
		if ns[i] >= ns[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

//双指针找最大值
func max1(start, len int) int {
	j := start + len - 1
	for start < j {
		if ns[start] > ns[j] {
			j--
		} else {
			start++
		}
	}
	return start
}
