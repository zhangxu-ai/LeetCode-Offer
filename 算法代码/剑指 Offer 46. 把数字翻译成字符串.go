package main

import (
	"fmt"
	"strconv"
)

//100,77
//和剑指10-1和10-2相似
//转换为string，从左到右遍历
func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	nums := strconv.Itoa(num)
	pre, pos := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1:i+1] >= "10" && nums[i-1:i+1] <= "25" {
			pre, pos = pos, pos+pre
		} else {
			pre = pos
		}
	}
	return pos
}

//从右向左求余，不需要转换为字符串,节约空间
//100,77
func translateNum1(num int) int {
	if num < 10 {
		return 1
	}
	nex, pos, x, y := 1, 1, 0, num%10
	for num != 0 {
		num /= 10
		x = num % 10
		if (x*10+y) >= 10 && (x*10+y) <= 25 {
			pos, nex = nex+pos, pos
		} else {
			nex = pos
		}
		y = x
	}
	return pos
}
func main() {
	num := 12058
	fmt.Println(translateNum1(num))
}
