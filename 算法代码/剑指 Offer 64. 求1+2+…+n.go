package main

/*
逻辑运算符的短路效应：
if(A && B)  // 若 A 为 false ，则 B 的判断不会执行（即短路），直接判定 A && B 为 false

if(A || B) // 若 A 为 true ，则 B 的判断不会执行（即短路），直接判定 A || B 为 true
*/
func sumNums(n int) int {
	res := 0

	var sum func(n int) bool
	sum = func(n int) bool {
		res += n
		return n > 1 && sum(n-1)
	}
	sum(n)
	return res
}
