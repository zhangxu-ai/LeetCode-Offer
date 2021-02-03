package main

import (
	"fmt"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	l := len(nums)
	if l == 0 {
		return ""
	}
	tmp := make([]string, l)
	for i := 0; i < l; i++ {
		tmp[i] = strconv.Itoa(nums[i])
	}
	tmp = quickSort(tmp)
	res := new(strings.Builder)
	for i := 0; i < l; i++ {
		res.WriteString(tmp[i])
	}
	return res.String()
}

func quickSort(strs []string) []string {
	l := len(strs)
	if l < 2 {
		return strs
	}
	base := strs[0]
	left, right := 0, l-1

	for left < right {
		for right > left && strs[right]+base >= base+strs[right] {
			right--
		}
		strs[right], strs[left] = strs[left], strs[right]
		for right > left && strs[left]+base <= base+strs[right] {
			left++
		}
		strs[right], strs[left] = strs[left], strs[right]
	}
	quickSort(strs[:left])
	quickSort(strs[left+1:])
	return strs
}

func main() {
	s := []int{3, 30, 34, 5, 9}
	fmt.Println(minNumber(s))
}
