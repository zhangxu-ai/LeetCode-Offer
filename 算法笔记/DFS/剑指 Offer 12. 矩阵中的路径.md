# 剑指 Offer 12. 矩阵中的路径

## 题目

请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

## 思考

标准的DFS，每进入到一个格子中都有可能需要再递归判断他相邻的格子；
在一趟遍历中，需要将已经路过的格子做好标记，防止二次进入
但是因为go的二维切片的参数可以被其他函数修改，所以在本趟改完后还需要再改回去才行。

## 代码

```go
//使用全局变量
var _board [][]byte
var _word string
func exist(board [][]byte, word string) bool {
	_board=board
	_word=word
	return _exist(0,0,0)
}

func _exist(i, j, k int) bool {
	if k>=len(_word){
		return true
	}
	if i<0||i>=len(_board)||j<0||j>=len(_board[0])||_board[i][j]!=_word[k]||_board[i][j]=='\n'{
		return false
	} else{
		_board[i][j]='\n'
		res:=_exist(i,j+1,k+1)||_exist(i+1,j,k+1)||_exist(i,j-1,k+1)||_exist(i-1,j,k+1)
		_board[i][j]=_word[k]
		return res
	}
}

```
