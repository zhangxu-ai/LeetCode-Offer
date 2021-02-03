package main

//DFS+剪枝
//99,78
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := helperIs(root)
	return res != -1
}
func helperIs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := helperIs(root.Left)
	//提前判断左，避免无谓的调用
	if l == -1 {
		return -1
	}
	r := helperIs(root.Right)
	if r == -1 || l-r > 1 || r-l > 1 {
		return -1
	}
	if l > r {
		return 1 + l
	} else {
		return 1 + r
	}
}
