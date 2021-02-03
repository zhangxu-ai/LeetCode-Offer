package main

var nodes []int
var kg int

func kthLargest(root *TreeNode, k int) int {
	nodes = make([]int, 0, k)
	kg = k
	helper(root)
	return nodes[k-1]
}

//从右边开始的中序遍历
func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Right)
	nodes = append(nodes, root.Val)
	if len(nodes) >= kg {
		return
	}
	helper(root.Left)
}
