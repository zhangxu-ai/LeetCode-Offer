package main

import "fmt"

//笨比法，直接超时
func lastRemaining(n int, m int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	for i := 1; i < n; i++ {
		if (m) <= len(nums) {
			nums = append(nums[m:], nums[:m-1]...)
		} else if m%len(nums) != 0 {
			nums = append(nums[m%len(nums):], nums[:m%len(nums)-1]...)
		} else {
			nums = append(nums[:len(nums)-1], nums[len(nums):]...)
		}
	}
	return nums[0]
}

//笨比暴力法，超时
func lastRemaining1(n int, m int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	index := 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < m-1; {
			for nums[index] == -1 {
				if index == n-1 {
					index = 0
				} else {
					index++
				}
			}
			j++
			if index == n-1 {
				index = 0
			} else {
				index++
			}
		}
		for nums[index] == -1 {
			if index == n-1 {
				index = 0
			} else {
				index++
			}
		}
		nums[index] = -1
		if index == n-1 {
			index = 0
		} else {
			index++
		}
	}
	for nums[index] == -1 {
		if index == n-1 {
			index = 0
		} else {
			index++
		}
	}
	return nums[index]
}

func main() {
	n, m := 5, 3
	fmt.Println(lastRemaining1(n, m))
}
