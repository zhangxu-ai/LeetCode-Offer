package main

import "fmt"

//原始解法，虽然使用二分进行了优化，但是还是暴力法
//48,92
func twoSum(nums []int, target int) []int {
	left := searchLeft1(nums, target)
	for i := left - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			t := nums[i] + nums[j]
			if t == target {
				return []int{nums[i], nums[j]}
			} else if t > target {
				break
			} else {
				continue
			}
		}
	}
	if nums[0] < 0 && nums[len(nums)-1] > target {
		for i := len(nums) - 1; i >= left; i-- {
			for j := 0; j <= i; j++ {
				if nums[j] > 0 {
					break
				}
				t := nums[i] + nums[j]
				if t == target {
					return []int{nums[i], nums[j]}
				} else if t > target {
					break
				} else {
					continue
				}
			}
		}
	}
	return []int{0, 0}
}

//寻找左边界
func searchLeft1(nums []int, target int) int {
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

//对每一个数，使用二分寻找其目标值，暴力法2
//50,96
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		l := BinarySearch(nums, target-nums[i])
		if l != -1 {
			return []int{nums[i], nums[l]}
		}
	}
	return nil
}
func BinarySearch(a []int, key int) int {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := (low + high) / 2
		if a[mid] < key {
			low = mid + 1
		} else if a[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

//对撞双指针
//88,92
//右上三角形排列
func twoSum2(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		t := target - nums[left]
		if t > nums[right] {
			left++
		} else if t < nums[right] {
			right--
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return nil
}

func main() {
	a := []int{16, 16, 18, 24, 30, 32}
	t := 48
	fmt.Println(twoSum1(a, t))
}
