package main

import "fmt"

func movingCount(m int, n int, k int) int {
	count := 0
	area := make([][]int, m)
	for i := 0; i < m; i++ {
		v := make([]int, n)
		area[i] = v
	}
	mdfs(area, 0, 0, k, &count)
	return count
}

//深度优先遍历
func mdfs(area [][]int, i, j, k int, count *int) {
	//越界检查
	if i >= len(area) || i < 0 || j >= len(area[0]) || j < 0 || area[i][j] == -1 {
		return
	}
	//k检查
	if (i/10 + i%10 + j/10 + j%10) > k {
		return
	}
	*count++
	area[i][j] = -1
	//只需要向下和右走就行，可以节省大量的内存空间（3.3->2.4）
	mdfs(area, i+1, j, k, count)
	mdfs(area, i, j+1, k, count)
}

//广度优先算法BFS

func main() {
	m, n, k := 2, 3, 1
	c := movingCount(m, n, k)
	fmt.Println(c)
}
