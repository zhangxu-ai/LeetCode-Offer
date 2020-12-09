package main

import "fmt"

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	f, s, r := 1, 1, 0
	for i := 2; i <= n; i++ {
		r = (f + s) % 1000000007
		f = s
		s = r
	}
	return r
}
func main() {
	n := []int{2, 7, 0}
	for _, i := range n {
		fmt.Println(numWays(i))
	}
}
