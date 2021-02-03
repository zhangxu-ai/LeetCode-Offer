package main

import (
	"fmt"
	"math"
)

//时间超过限制
//甚至连电子垃圾都算不上
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	r := 0
	//nums = append(nums, "0")
	cur := 10
	for i := 10; i <= n; {
		d := 0
		for j := cur; j != 0; j /= 10 {
			d++
		}
		t := cur
		for k := d; k > 0; k-- {
			p := int(math.Pow10(k - 1))
			r = t / p
			t %= p
			if i == n {
				return r
			}
			i++
		}
		cur++
	}
	return r
}

//100,70
func findNthDigit1(n int) int {
	if n < 10 {
		return n
	}
	var (
		digit      = 1
		numLen     = 1
		start, end = 0, 9 * digit
	)
	for n >= end {
		digit *= 10
		numLen++
		start = end + 1
		end = start + numLen*digit*9 - 1
	}
	p1 := (n-start)/numLen + digit
	t := 0
	for i := (n - start) % numLen; i >= 0; i-- {
		t = p1 / digit
		p1 %= digit
		digit /= 10
	}
	return t
}

func main() {
	n := 0
	fmt.Println(findNthDigit1(n))

}
