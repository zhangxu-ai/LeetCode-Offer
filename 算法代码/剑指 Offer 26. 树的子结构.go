package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

//使用暴力搜索的方法 ，相比递归，时间花费少很多
func isSubStructure1(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	nodes := []*TreeNode{A}
	l := 0
	for l <= len(nodes)-1 {
		if nodes[l].Left != nil {
			nodes = append(nodes, nodes[l].Left)
		}
		if nodes[l].Right != nil {
			nodes = append(nodes, nodes[l].Right)
		}
		if recur(nodes[l], B) {
			return true
		}
		l++
	}
	return false
}
