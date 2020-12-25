package main

//动态规划+数学推导
//100,95
//丑数的递推性质： 丑数只包含因子 2, 3, 52,3,5 ，
//因此有 “丑数 == 某较小丑数 × 某因子”
//例如：10 = 5 * 2
func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0
	min := func(a, b, c int) int {
		t := a
		if t > b {
			t = b
		}
		if t > c {
			t = c
		}
		return t
	}
	for i := 1; i < n; i++ {
		l, m, n := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(l, m, n)
		if dp[i] == l {
			a++
		}
		if dp[i] == m {
			b++
		}
		if dp[i] == n {
			c++
		}
	}
	return dp[n-1]
}
