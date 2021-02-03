package main

import (
	"fmt"
	"math"
)

//100,97
func verifyPostorder(postorder []int) bool {
	var recur func(i, j int) bool
	recur = func(i, j int) bool {
		//如果left==right，就一个节点不需要判断了，如果left>right说明没有节点，
		//也不用再看了,否则就要继续往下判断
		if i >= j {
			return true
		}
		//遍历后序遍历的 [i, j] 区间元素，寻找 第一个大于根节点 的节点，索引记为 m 。
		p, m := i, 0
		for postorder[p] < postorder[j] {
			p++
		}
		m = p
		//之后的所有节点如果比root大，则会一直便利到p==j.
		for postorder[p] > postorder[j] {
			p++
		}
		//如果p!=j,则说明后面（右子树区间）出现了比root小的值，即为非法的后续遍历顺序
		return p == j && recur(i, m-1) && recur(m, j-1)
	}
	return recur(0, len(postorder)-1)
}

//100,67.99
func verifyPostorder1(postorder []int) bool {
	var stack []int
	//root一开始为无穷大，
	root := math.MaxInt64
	//倒序遍历此数组，
	for i := len(postorder) - 1; i >= 0; i-- {
		//第一次运行到这一定会跳过，因为root是最大值
		//第二次时还是最大值，但是在后面有可能会将倒数第二个值赋给root
		if postorder[i] > root {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
			l := len(stack) - 1
			root = stack[l]
			stack = stack[:l]
		}
		stack = append(stack, postorder[i])
	}
	return true
}

func main() {
	a := []int{1, 3, 2, 6, 5}
	fmt.Println(verifyPostorder(a))
}
