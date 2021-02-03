package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	//默认在最后一行
	i, j, k := len(matrix)-1, len(matrix[0])-1, 0
	//确定截止行
	//从最后一行开始，找到某一行，此行的第一个元素比目标值要小，说明该行之前可能有目标值
	for ; i >= 0; i-- {
		if matrix[i][0] == target {
			return true
		}
		if matrix[i][0] < target {
			break
		}
	}
	//找到每一行的最后一个值，确定开始行
	for r := 0; r <= i; r++ {
		if matrix[r][j] == target {
			return true
		}
		if matrix[r][j] > target {
			k = r
			break
		}
	}
	//使用二分查找遍历每一行
	for ; k <= i; k++ {
		if binSearch(matrix[k], target) {
			return true
		}
	}
	return false
}

//二分查找
func binSearch(a []int, t int) bool {
	i, j := 0, len(a)-1
	for i <= j {
		mid := (i + j) / 2
		if t == a[mid] {
			return true
		} else if t > a[mid] {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return false
}

/*
从矩阵 matrix 左下角元素（索引设为 (i, j) ）开始遍历，并与目标值对比：
当 matrix[i][j] > target 时，执行 i-- ，即消去第 i 行元素；
当 matrix[i][j] < target 时，执行 j++ ，即消去第 j 列元素；
当 matrix[i][j] = target 时，返回 truetrue ，代表找到目标值。
若行索引或列索引越界，则代表矩阵中无目标值，返回 falsefalse 。
*/
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := len(matrix)-1, 0
	for i >= 0 && j <= len(matrix[0])-1 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return false
}

func main() {
	m := [][]int{{-5, 3}}
	t := 3
	fmt.Println(findNumberIn2DArray1(m, t))
}
