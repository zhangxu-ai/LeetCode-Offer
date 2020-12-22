package stl

import "fmt"

func GetNext(s string) []int {
	next := make([]int, len(s)+1)
	next[0] = -1
	var j = -1
	for i := 0; i < len(s); {
		//j==-1说明第一次遍历或者在后续遍历中
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func IndexKmp(s string, t string) int {
	if s == "" || t == "" || len(s) < len(t) {
		return -1
	}
	next := GetNext(t)
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if j == -1 || s[i] == t[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(t) {
		return i - len(t)
	}
	return -1
}

func main() {
	s := "1234512345"
	t := "3451"
	fmt.Println(IndexKmp(s, t))
}
