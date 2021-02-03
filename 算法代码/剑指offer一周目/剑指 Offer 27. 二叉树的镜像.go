package main

//递归做法
//时间100%,内存61%
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//res:=&TreeNode{
	//	Val:   root.Val,
	//	Left: mirrorTree(root.Right),
	//	Right: mirrorTree(root.Left),
	//}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

//广度优先
//时间100%,内存18%

func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	l := 0
	for l <= len(nodes)-1 {
		if nodes[l].Right != nil {
			nodes = append(nodes, nodes[l].Right)
		}
		if nodes[l].Left != nil {
			nodes = append(nodes, nodes[l].Left)
		}
		nodes[l].Left, nodes[l].Right = nodes[l].Right, nodes[l].Left
		l++
	}
	return root
}
