package main

import (
	"fmt"
)

func printNumbers(n int) []int {
	m := int(myPowP(10, n))
	arr := make([]int, m-1)

	for i := 1; i < m; i++ {
		arr[i-1] = i
	}
	return arr
}

//使用
func myPowP(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	var r = 1.0
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n > 0 {
		if (n & 1) == 1 {
			r *= x
		}
		x *= x
		n >>= 1
	}
	return r
}

////考虑大数问题
//func printNumbers1(n int) []int {
//
//}

func main() {
	n := 2
	fmt.Println(printNumbers(n))
}
