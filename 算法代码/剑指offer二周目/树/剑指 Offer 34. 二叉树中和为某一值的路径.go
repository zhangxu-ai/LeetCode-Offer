package main

var (
	rsp  [][]int
	path []int
)

func pathSum(root *TreeNode, sum int) [][]int {
	rsp = nil
	if root == nil {
		return rsp
	}
	_pathSum(root, 0, sum)
	return rsp
}

func _pathSum(root *TreeNode, now, target int) {
	now = now + root.Val
	path = append(path, root.Val)
	if root.Left != nil {
		_pathSum(root.Left, now, target)
	}
	if root.Right != nil {
		//pathRight:=make([]int,len(path))
		//copy(pathRight, path)
		_pathSum(root.Right, now, target)
	}
	if root.Left == nil && root.Right == nil && now == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		rsp = append(rsp, tmp)
	}
	path = path[:len(path)-1]
}
