package main

import "fmt"

//动态规划
//73,63
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l <= 1 {
		return l
	}
	//字符的 ASCII 码范围为 00 ~ 127127 ，
	//哈希表最多使用 O(128) = O(1)O(128)=O(1) 大小的额外空间。
	positions := make(map[byte]int, 128)
	positions[s[0]] = 0
	max, pre := 1, 1
	for i := 1; i < l; i++ {
		ind, ok := positions[s[i]]
		if ok && i-ind <= pre {
			pre = i - ind
		} else {
			pre += 1
		}
		//if !ok{
		//	pre+=1
		//}else{
		//	if i-ind>pre{
		//		pre+=1
		//	}else{
		//
		//	}
		//}
		//更新各字符出现的最后一个位置
		positions[s[i]] = i
		if pre > max {
			max = pre
		}
	}
	return max
}

//动态规划，也叫滑动窗口
func lengthOfLongestSubstring1(s string) int {
	l := len(s)
	if l <= 1 {
		return l
	}
	left, right, max := 0, 0, 0
	positions := make(map[byte]int, 128)
	for ; right < l; right++ {
		ind, ok := positions[s[right]]
		//视情况更新左边界，其他情况不变
		if ok && ind >= left {
			left = ind + 1
		}
		positions[s[right]] = right
		if max < right-left+1 {
			max = right - left + 1
		}
	}
	return max
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring1(s))
}
