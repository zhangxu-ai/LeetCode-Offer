package main

//递归即可
//89,75
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := maxDepth(root.Left), maxDepth(root.Right)
	if l > r {
		return 1 + l
	} else {
		return 1 + r
	}
}

//BFS
//89,5
type nodeWithHeight struct {
	root   *TreeNode
	height int
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*nodeWithHeight{{
		root:   root,
		height: 1,
	}}
	res := new(nodeWithHeight)
	for len(stack) != 0 {
		if stack[0].root.Left != nil {
			stack = append(stack, &nodeWithHeight{
				root:   stack[0].root.Left,
				height: stack[0].height + 1,
			})
		}
		if stack[0].root.Right != nil {
			stack = append(stack, &nodeWithHeight{
				root:   stack[0].root.Right,
				height: stack[0].height + 1,
			})
		}
		res = stack[0]
		stack = stack[1:]
	}
	return res.height
}

//BFS2
//看别人答案发现不需要使用结构体
//89,96
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res, que := 0, []*TreeNode{root}
	for len(que) > 0 {
		res++
		n := len(que)
		for i := 0; i < n; i++ {
			node := que[0]
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			que = que[1:]
		}
	}
	return res
}
