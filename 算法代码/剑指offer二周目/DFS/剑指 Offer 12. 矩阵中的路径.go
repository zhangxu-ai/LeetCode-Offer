package main

import "fmt"

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

func main() {
	b:=[][]byte{
		{'a','b','c','e'},
		{'s','f','c','s'},
		{'a','d','e','e'},
	}
	w:="abcced"
	fmt.Println(exist(b,w))
}
