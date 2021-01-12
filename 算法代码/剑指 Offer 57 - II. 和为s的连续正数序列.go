package main

import (
	"fmt"
	"math"
)

//推导数学公式
//起始数为n,假设到n+k结束，则他们和目标值t的关系为
//n+n+1+...+n+k=t
//(k+1)n+(k*(k+1))/2=t
//求k的正整数根
//100,100
func findContinuousSequence(target int) [][]int {
	var res [][]int
	for i := 1; i <= target/2; i++ {
		b := 2*i + 1
		d := b*b - 8*(i-target)
		if d < 0 {
			continue
		}
		dt := math.Sqrt(float64(d))
		l, r := math.Modf(dt)
		if r != 0 || (int(l)-b)%2 != 0 {
			continue
		} else {
			k := (int(l) - b) / 2
			tmp := make([]int, 0, k+1)
			for j := i; j <= i+k; j++ {
				tmp = append(tmp, j)
			}
			res = append(res, tmp)
		}

	}
	return res
}

//滑动窗口法
//100,100
func findContinuousSequence1(target int) [][]int {
	var res [][]int
	left, right := 1, 1
	for ; left <= target/2; left++ {
		t := ((right - left) + 1) * (left + right) / 2
		for t < target {
			right++
			t += right
		}
		if t == target {
			tmp := make([]int, 0, (right-left)+1)
			for j := left; j <= right; j++ {
				tmp = append(tmp, j)
			}
			res = append(res, tmp)
		} else {
			continue
		}
	}
	return res
}
func main() {
	fmt.Println(findContinuousSequence1(9))
}
