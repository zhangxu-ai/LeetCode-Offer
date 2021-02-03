package main

import "fmt"

//82,87
func validateStackSequences(pushed []int, popped []int) bool {
	l := len(pushed)
	if l == 0 {
		return true
	}
	help := make([]int, 0, l)
	for i, j := 0, 0; j < l; i++ {
		//先检查出栈操作
		tl := len(help)
		for tl > 0 && help[tl-1] == popped[j] {
			help = help[:tl-1]
			j++
			tl--
		}
		//再入栈
		if i < l {
			t := pushed[i]
			help = append(help, t)
			tl++
			if t == popped[j] {
				help = help[:tl-1]
				j++
			}
		} else { //说明入栈都入完了
			if j >= l {
				return true
			} else if help[tl-1] == popped[j] {
				help = help[:tl-1]
				j++
			} else {
				return false
			}
		}
	}
	return len(help) == 0
}

//82,91
func validateStackSequences1(pushed []int, popped []int) bool {
	l := len(pushed)
	if l == 0 {
		return true
	}
	help := make([]int, 0, l)
	for i, j := 0, 0; i < l; i++ {
		help = append(help, pushed[i])
		for k := len(help); k > 0 && help[k-1] == popped[j]; {
			help = help[:k-1]
			k--
			j++
		}
	}
	return len(help) == 0
}
func main() {
	a := []int{2, 3, 0, 1}
	b := []int{0, 3, 2, 1}
	fmt.Println(validateStackSequences1(a, b))
}
