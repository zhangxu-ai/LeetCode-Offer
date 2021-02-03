package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}
	var res = make([]int, 0, m*n)

	i, j, ul, dl, lc, rc := 0, 0, 0, 0, 0, 0
	for ul+dl < m && lc+rc < n {
		//遍历第一步:左->右
		for ; lc+rc < n && j < n-rc; j++ {
			res = append(res, matrix[i][j])
		}
		j--
		i++
		ul++
		//遍历第二步：上—>下
		for ; ul+dl < m && i < m-dl; i++ {
			res = append(res, matrix[i][j])
		}
		i--
		j--
		rc++
		//遍历第三步:右->左
		for ; ul+dl < m && lc+rc < n && j >= lc; j-- {
			res = append(res, matrix[i][j])
		}
		j++
		i--
		dl++
		//遍历第四步：下—>上
		for ; lc+rc < n && ul+dl < m && i >= ul; i-- {
			res = append(res, matrix[i][j])
		}
		i++
		j++
		lc++
	}
	return res
}

func main() {
	a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(a))
}
