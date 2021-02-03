package main

//DFS
//90,84
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	dfsP(root, sum, []int{}, &res)
	return res
}

func dfsP(root *TreeNode, sum int, way []int, res *[][]int) {
	if root == nil {
		return
	}
	way = append(way, root.Val)
	sum -= root.Val
	//如果到了叶子节点且达成了目标
	if sum == 0 && root.Left == nil && root.Right == nil {
		var tmp = make([]int, len(way))
		copy(tmp, way)
		*res = append(*res, tmp)
	}

	//递归左子树
	dfsP(root.Left, sum, way, res)
	//递归右子树
	dfsP(root.Right, sum, way, res)
}
