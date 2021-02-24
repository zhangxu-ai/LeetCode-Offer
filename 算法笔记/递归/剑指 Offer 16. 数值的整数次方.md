# 剑指 Offer 16. 数值的整数次方

## 题目

实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

## 思考

n为偶数时： x^n=x^(n/2)*x^(n/2)
n为奇数时： x^n=x^(n/2)*x^(n/2) * x

**快速幂**：根据上面的公式，可以使用递归求，时间降低到 O(logn)

## 代码

```go
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := pow(x, n/2)
	if n%2 == 1 {
		return half * half * x
	} else {
		return half * half
	}
}

```
