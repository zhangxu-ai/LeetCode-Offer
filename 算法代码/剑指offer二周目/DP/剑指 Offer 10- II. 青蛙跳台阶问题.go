package main

//递归版本，超出时间限制
//func numWays(n int) int {
//	if n==0||n==1{
//		return 1
//	}
//	return (numWays(n-1)%1000000007+numWays(n-2)%1000000007)%1000000007
//}

//DP版本
func numWays(n int) int {
	if n==0||n==1{
		return 1
	}
	//第一版，使用数组记录每个情况
	//dp:=make([]int,n+1)
	//dp[0]=1
	//dp[1]=1
	//for i := 2; i <=n; i++ {
	//	dp[i]=(dp[i-1]+dp[i-2])%1000000007
	//}
	//return dp[n]
	//我们发现，每个情况只和他前面的一种情况有关，
	//不需要再往前的值，所以可以只使用2个变量即可
	two,one:=1,1
	res:=0
	for i := 2; i <=n; i++ {
		res=(two+one)%1000000007
		two=one
		one=res
	}
	return res
}
