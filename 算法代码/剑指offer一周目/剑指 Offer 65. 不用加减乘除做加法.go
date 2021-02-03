package main

//数电的加法器
//使用位运算
//a+b的不考虑进位的结果是：a^b
//进位是：a&b<<1
//结果是进位+和，还是需要加号，所以需要使用for 循环或者递归
//跌代
func add(a int, b int) int {
	if a == 0 {
		return b
	}
	for b != 0 {
		c := (a & b) << 1
		a ^= b
		b = c
	}
	return a
}

//递归解决
func add1(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	c := (a & b) << 1
	a ^= b
	b = c
	return add1(a, b)
}
