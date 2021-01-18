package main

import "fmt"

//自己想的，dp+dfs
//100,36
//虽然花了很长时间，但是很爽
func dicesProbability(n int) []float64 {
	//最小为1*n，最大为6*n，总共 5*n+1种情况，
	allSum := 5*n + 1
	res := make([]float64, 0, allSum)
	//先考虑只有一个骰子的情况
	if n == 1 {
		for i := 1; i <= 6; i++ {
			res = append(res, 1.0/6)
		}
		return res
	}
	//n个色子能出现的所有情况的个数6的n次方
	all := myPowInt(6, n)
	//第一种情况一定是n个1，只有1种，可以先确定下来
	res = append(res, 1/all)
	//dp的k1为骰子个数，k2为目标值，v为k1个骰子出现总和为k2的情况的个数
	dp := make(map[int]map[int]int, n)
	//1个骰子出现1-6的情况都是1种
	dp[1] = make(map[int]int, 3)
	for i := 1; i <= 3; i++ {
		dp[1][i] = 1
	}
	//每个骰子数情况中，和出现的情况数目都是对称的，
	//所以只要存前一半的情况数目就行，节省空间
	//总的和的情况分为偶数和奇数，为了方便，多取一个，
	for i := 2; i <= n; i++ {
		tmp := make(map[int]int, 3*i)
		dp[i] = tmp
	}
	//遍历所有和的可能情况
	//因为1*n的情况我们上面已经考虑过了，所以从1*n+1开始遍历
	//通过找规律可知，res是对称分布的，所以只要算出前一半的概率即可
	for i := n + 1; i <= 7*n/2; i++ {
		res = append(res, float64(helperDp(&dp, n, i))/all)
	}
	start := 0
	//将后一半放上去
	if 7*n%2 == 0 {
		start = 7*n/2 - n - 1
	} else {
		start = 7*n/2 - n
	}
	for i := start; i >= 0; i-- {
		res = append(res, res[i])
	}
	return res
}

//递归计算
func helperDp(dp *map[int]map[int]int, n, target int) int {
	//情况对称，取左半边的情况数即可
	if target > 7*n/2 {
		target = 7*n - target
	}
	//如果dp中有存储过，则直接返回
	if v, ok := (*dp)[n][target]; ok {
		return v
	}
	//否则计算
	//确定第一个骰子能取的最大个数，
	//比如如果n=2,target=5,那么第一个骰子最大只能取到4，
	//因为第二个最小是1
	max := target - (n - 1)
	if target > ((n - 1) + 6) {
		max = 6
	}
	tmp := 0
	//遍历第一个骰子能取得情况下，
	//剩余n-1个骰子所需要取到的值的情况数目
	for i := 1; i <= max; i++ {
		tmp += helperDp(dp, n-1, target-i)
	}
	//更新dp
	(*dp)[n][target] = tmp
	return tmp
}
func myPowInt(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	var r = 1.0
	for n > 0 {
		if (n & 1) == 1 {
			r *= x
		}
		x *= x
		n >>= 1
	}
	return r
}

func main() {
	t := 2
	fmt.Println(dicesProbability(t))
}
