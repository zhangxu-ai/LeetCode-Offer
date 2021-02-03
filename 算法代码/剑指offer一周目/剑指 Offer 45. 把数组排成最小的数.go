package main

import (
	"fmt"
	"sort"
	"strings"
)

//29,58
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		s1 := fmt.Sprintf("%d%d", nums[i], nums[j])
		s2 := fmt.Sprintf("%d%d", nums[j], nums[i])
		if s1 < s2 {
			return true
		}
		return false
	})
	var res strings.Builder
	for i := 0; i < len(nums); i++ {
		res.WriteString(fmt.Sprintf("%d", nums[i]))
	}
	return res.String()
	//res:=""
	//for i := 0; i < len(nums); i++ {
	//	res+=strconv.Itoa(nums[i])
	//}
	//return res
}
