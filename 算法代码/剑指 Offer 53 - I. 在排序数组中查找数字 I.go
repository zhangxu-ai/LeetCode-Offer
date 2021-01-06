package main

import "fmt"

//最简单解法
//99,62
func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 || target < nums[0] || target > nums[l-1] {
		return 0
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			sum++
		}
	}
	return sum
}

//排序数组，使用二分法查找左边得边界，理论上会快
//也可以再次寻找右边界
//89,100
func search1(nums []int, target int) int {
	l := len(nums)
	if l == 0 || target < nums[0] || target > nums[l-1] {
		return 0
	}
	left := searchLeft(nums, target)
	//检测left越界情况
	if left >= l || nums[left] != target {
		return 0
	}
	//再向右遍历
	sum := 0
	for i := left; i < l; i++ {
		if nums[i] > target {
			break
		}
		sum++
	}
	return sum
	////找到target+1的左边界，即为target的右边界+1
	////89,62
	//right:=searchLeft(nums,target+1)
	//return right-left
}

func searchLeft(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	//二分查左侧边界
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}

func main() {
	t := 8
	a := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search1(a, t))
}
