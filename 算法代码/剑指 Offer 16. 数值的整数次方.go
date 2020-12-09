package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
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
func myPow1(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	r, c, m := x, 1, n/2
	for m != 0 {
		r *= r
		m /= 2
		c *= 2
	}
	return r * myPow1(x, n-c)
}
func main() {
	x := 2.0
	n := 10
	fmt.Printf("%f", myPow1(x, n))
}
