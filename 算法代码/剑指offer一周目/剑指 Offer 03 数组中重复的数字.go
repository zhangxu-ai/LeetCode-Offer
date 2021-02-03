package main

import (
	"fmt"
	"sort"
)

//使用hash表进行判断
//时间复杂度：O(n)
//空间复杂度：O(n)
func findRepeatNumber1(nums []int) int {
	nums1 := map[int]int{}
	for _, v := range nums {
		if _, ok := nums1[v]; ok {
			return v
		} else {
			nums1[v] = v
		}
	}
	return -1
}

//sort后判断是否有相同的
//时间复杂度大于O(n)
//空间复杂度小于O(n)
func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for i := range nums {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

//原地置换法
//假设理想情况：index=arr[index],
//如果不等于，则可能有重复，将此value与arr[value]比较，如果相等，则说明重复了
//如果不相等，则进行替换，一直到index=arr[index]为止，再进行下一个index的比较
func findRepeatNumber3(nums []int) int {
	t := 0
	for i := 0; i < len(nums); i++ {
		//如果当前index不等于当前的值
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			t = nums[i]
			nums[i] = nums[t]
			nums[t] = t
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 2, 5, 6, 8}
	fmt.Println(findRepeatNumber3(nums))
}
