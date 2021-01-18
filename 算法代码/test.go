package main

import "fmt"

func main() {
	//fmt.Println(1/6.0)
	dp := make(map[int]map[int]int, 1)
	if dp[0] == nil {
		fmt.Println(1)
	}
}
