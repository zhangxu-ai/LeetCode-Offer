package main

import "fmt"

//动态规划
//83,38
func maxValue(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = grid[i-1][j-1] + max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	//for _, ints := range dp {
	//	fmt.Println(ints)
	//}
	return dp[m-1][n-1]
}

//动态规划2，原地变换，不使用dp数组
//83,100
func maxValue1(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] = grid[i-1][j] + grid[i][j]
				} else {
					grid[i][j] = grid[i][j-1] + grid[i][j]
				}
			}
		}
	}
	//for _, ints := range dp {
	//	fmt.Println(ints)
	//}
	return grid[m-1][n-1]
}

//动态规划1的改进
//dp比原数组多1行1列，可以不考虑边界问题
func maxValue11(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := 0; i <= len(grid); i++ {
		dp[i] = make([]int, n+1)
	}
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[i][j] = grid[i-1][j-1] + max(dp[i][j-1], dp[i-1][j])
		}
	}
	return dp[m][n]
}

func main() {
	a := [][]int{{3}}
	fmt.Println(maxValue1(a))
}
