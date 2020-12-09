package main

import "fmt"

func exist(board [][]byte, word string) bool {
	//循环数组，每次都从word的第一个字符开始匹配
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

//深度优先遍历
func dfs(board [][]byte, word string, i, j, k int) bool {
	//如果越界或者不等，返回false
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	//如果搜索到头了
	if k == len(word)-1 {
		return true
	}
	//以下情况则是当前byte正好是word[k]
	k++
	//标记当前元素已经被访问过了
	board[i][j] = ' '
	res := dfs(board, word, i-1, j, k) || //上
		dfs(board, word, i+1, j, k) || //下
		dfs(board, word, i, j-1, k) || //左
		dfs(board, word, i, j+1, k) //右
	//还原数组//别忘了k被加1了
	board[i][j] = word[k-1]
	return res
}

func main() {
	a := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'F'}}
	word := "ABCCED"
	fmt.Println(exist(a, word))
}
