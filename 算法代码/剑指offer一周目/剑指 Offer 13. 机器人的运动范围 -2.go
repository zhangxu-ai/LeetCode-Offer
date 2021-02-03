package main

import "fmt"

func movingCount1(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		v := make([]bool, n)
		visited[i] = v
	}
	count := 0
	//使用go的切片模拟队列
	var queue [][]int
	queue = append(queue, []int{0, 0})
	for len(queue) > 0 {
		i := queue[0][0]
		j := queue[0][1]
		queue = append(queue[:0], queue[1:]...)

		if i >= m || i < 0 || j >= n || j < 0 || visited[i][j] || ((i/10 + i%10 + j/10 + j%10) > k) {
			continue
		}
		visited[i][j] = true
		count++
		queue = append(queue, []int{i + 1, j})
		queue = append(queue, []int{i, j + 1})
	}
	return count
}
func main() {
	m, n, k := 1, 2, 1
	c := movingCount1(m, n, k)
	fmt.Println(c)
}
