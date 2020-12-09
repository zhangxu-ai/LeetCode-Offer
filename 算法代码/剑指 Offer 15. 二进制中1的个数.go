package main

import (
	"fmt"
	"strings"
)

func hammingWeight(num uint32) int {
	r := 0
	for num != 0 {
		if num&1 == 1 {
			r++
		}
		num = num >> 1
		fmt.Printf("%32b\r\n", num)
	}
	return r
}

//使用fmt函数将num转换为字符串，再调用strings包自带的函数进行转换
func hammingWeight1(num uint32) int {
	return strings.Count(fmt.Sprintf("%b", num), "1")
}
