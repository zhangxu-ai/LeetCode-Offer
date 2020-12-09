package main

import (
	"fmt"
)

//和剪绳子1不同的是，有大数，不能使用math包的求幂运算，会有精度损失
//改用循环求幂法
func p3n(p int) int {
	r := 1
	for i := 0; i < p-1; i++ {
		r = r * 3 % 1000000007
	}
	return r
}

func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	p := n / 3
	d := n % 3
	r := p3n(p)
	if d == 2 {
		r *= 3 * 2
	} else if d == 1 {
		r *= 2 * 2
	} else {
		r *= 3
	}
	return r % 1000000007
}

func main() {
	n := 500
	fmt.Println(cuttingRope2(n))
}
