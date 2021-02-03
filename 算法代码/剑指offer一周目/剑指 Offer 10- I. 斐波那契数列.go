package main

import "fmt"

//递归求解，时间复杂度最高，不通过
func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}
	return (fib(n-1) + fib(n-2)) % 1000000007
}

//递推求解,动态规划
func fib1(n int) int {
	f := make([]int, n+2)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		//注意，每次计算后就要取模，否者会溢出
		//取模后相加再取模，模不变
		f[i] = (f[i-2] + f[i-1]) % 1000000007
	}
	return f[n]
}

//递推优化：
//因为每一步只需要前2个数，不需要把每个结果都存起来，所以使用变量接收参数
func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	//f:a(n-2)
	//s:a(n-1)
	//r:a(n)
	f, s, r := 0, 1, 0
	for i := 2; i <= n; i++ {
		r = (f + s) % 1000000007
		f = s
		s = r
	}
	return r
}

//递归实现2，只是把每一个结果暂存起来，和递归类似
func fib3(n int) int {
	f := make([]int, n+2)
	f[0] = 0
	f[1] = 1
	return fib3_helper(n, f)
}
func fib3_helper(n int, f []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if f[n] != 0 {
		return f[n]
	}
	r := (fib3_helper(n-2, f) + fib3_helper(n-1, f)) % 1000000007
	f[n] = r
	return r
}
func main() {
	n := 5
	fmt.Println(fib2(n))
}
