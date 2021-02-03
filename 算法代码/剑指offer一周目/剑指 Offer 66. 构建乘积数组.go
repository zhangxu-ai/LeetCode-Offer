package main

import "fmt"

//笨比法
//10,5
func constructArr(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}
	pre, tail, res := make([]int, l), make([]int, l), make([]int, l)
	pre[0] = 1
	tail[l-1] = 1
	for i := 1; i < l; i++ {
		pre[i] = pre[i-1] * a[i-1]
		tail[l-1-i] = tail[l-i] * a[l-i]
	}
	for i := 0; i < l; i++ {
		res[i] = pre[i] * tail[i]
	}
	return res
}

//大佬优化：
//计算pre的时候，可以不使用额外的数组来存，直接放到b里面就行
//计算tail的时候，因为是依次往前算，所以可以使用一个常数来存，不需要tail数组
func constructArr1(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}
	res, tmp := make([]int, l), 1
	res[0] = 1
	for i := 1; i < l; i++ {
		res[i] = res[i-1] * a[i-1]
	}
	for i := l - 2; i >= 0; i-- {
		tmp *= a[i+1]
		res[i] = res[i] * tmp
	}
	return res
}
func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr1(a))
}
