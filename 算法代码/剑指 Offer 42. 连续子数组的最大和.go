package main

import "fmt"

//第一种思想：看似动态规划，其实还是暴力遍历的方法，自作聪明了
//提交结果超时 O(N2)
func maxSubArray0(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	var sums []int
	max := -25536
	for i := l - 1; i >= 0; i-- {
		for j := len(sums) - 1; j >= 0; j-- {
			sums[j] = sums[j] + nums[i]
			if sums[j] > max {
				max = sums[j]
			}
		}
		sums = append(sums, nums[i])
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//真*动态规划
//从头遍历，维护一个一维数组dp[],dp[i]代表以i结尾的连续子数组的最大值
//如果dp[i-1]>=0,则i要加上前一个数
//如果dp[i-1]<0,则nums[i]本身就是最大的子数组，前面二点元素都是拖后腿的
//48,95
func maxSubArray1(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}

	pre, max := nums[0], nums[0]
	for i := 1; i < l; i++ {
		if pre < 0 {
			pre = nums[i]
		} else {
			pre = pre + nums[i]
		}
		if pre >= max {
			max = pre
		}
	}
	return max
}

//动态规划，原地修改
//48,35
func maxSubArray2(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	a := []int{-1}
	fmt.Println(maxSubArray1(a))
}
