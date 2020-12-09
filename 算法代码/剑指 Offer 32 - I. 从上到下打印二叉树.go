package main

//BFS
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	tmp := []*TreeNode{root}
	res := []int{root.Val}
	for i := 0; i < len(tmp); i++ {
		if tmp[i].Left != nil {
			tmp = append(tmp, tmp[i].Left)
			res = append(res, tmp[i].Left.Val)
		}
		if tmp[i].Right != nil {
			tmp = append(tmp, tmp[i].Right)
			res = append(res, tmp[i].Right.Val)
		}
	}
	return res
}

//DFS
