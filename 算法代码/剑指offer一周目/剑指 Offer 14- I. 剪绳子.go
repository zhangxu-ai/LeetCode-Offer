package main

import (
	"fmt"
	"math"
	"sort"
)

//动态规划法
var store = map[int]int{1: 1}

func cuttingRope(n int) int {
	if r, ok := store[n]; ok {
		return r
	}
	res := make([]int, n/2)
	//从1到n的一半，算出每次的结果
	for i := 1; i <= n/2; i++ {
		m := cuttingRope(i)
		l := cuttingRope(n - i)
		if m < i {
			m = i
		}
		if l < n-i {
			l = n - i
		}
		res[i-1] = m * l
	}
	sort.Ints(res)
	store[n] = res[n/2-1]
	return res[n/2-1]
}

//使用数学分析法做，内存使用减少了
func cuttingRope1(n int) int {
	if n <= 3 {
		return n - 1
	}
	p := n / 3
	d := n % 3
	r := math.Pow(3, float64(p-1))
	if d == 2 {
		r *= 3 * 2
	} else if d == 1 {
		r *= 2 * 2
	} else {
		r *= 3
	}
	return int(r)
}

func main() {
	n := 58
	r := cuttingRope1(n)%1e9 + 7
	fmt.Println(r)
}
