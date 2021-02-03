package main

import "sort"

//100,82
//艹了，一副牌还带3个王的
func isStraight(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			joker++
		}
	}
	for i := joker; i < 4; {
		t := nums[i+1] - nums[i]
		if t == 1 {
			i++
		} else if t > (joker+1) || t == 0 {
			return false
		} else {
			joker -= t - 1
			i += t - 1
		}
	}
	return true
}

//大佬简化后的逻辑
//流下了不学无术的泪水
func isStraight1(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			//重复的情况
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}
