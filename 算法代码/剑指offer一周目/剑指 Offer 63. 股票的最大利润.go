package main

//笨比遍历
func maxProfit(prices []int) int {
	r := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > r {
				r = prices[j] - prices[i]
			}
		}
	}
	return r
}

//想到双指针发->滑动窗口法
//想岔了😭
//func maxProfit1(prices []int) int {
//	buy,sell:=0,1
//	for i := 2; i < len(prices); i++ {
//		//当前是否亏钱，如果亏钱，说明一定买贵了，
//		//所以把买的日期后延，相应的卖的日期也要后延
//		if prices[sell]<=prices[buy]{
//			buy++
//			sell++
//		}else{ //如果在赚钱，则看今天是涨还是跌
//			//如果跌的不多，则继续持有
//			if prices[i]<prices[sell]{
//				if prices[i]>=prices[buy]{
//
//				}
//			}
//		}
//	}
//
//}

//动态规划
//100,100
func maxProfit1(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	//dp:=make([]int,l)
	pre, min := 0, prices[0]
	for i := 1; i < l; i++ {
		if prices[i]-min > dp[i-1] {
			pre = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return pre
}
