package main

//100,84
func levelOrderIII(root *TreeNode) [][]int {
	var (
		nodes    []*TreeNode //本次遍历的nodes
		res      [][]int     //总的返回值
		tmp      []int       //缓存的本层值
		tmpNodes []*TreeNode //缓存的下一层的nodes
		level    = 1         //层次，从1开始，即root层
	)
	if root != nil {
		nodes = append(nodes, root)
	}
	for len(nodes) > 0 {
		for i := len(nodes); i > 0; i-- {
			node := nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			tmp = append(tmp, node.Val)
			//如果是偶数层，则先入栈右节点，再入左节点
			if level%2 == 0 {
				if node.Right != nil {
					tmpNodes = append(tmpNodes, node.Right)
				}
				if node.Left != nil {
					tmpNodes = append(tmpNodes, node.Left)
				}
			} else {
				if node.Left != nil {
					tmpNodes = append(tmpNodes, node.Left)
				}
				if node.Right != nil {
					tmpNodes = append(tmpNodes, node.Right)
				}
			}
		}
		res = append(res, tmp)
		//进入下一层的准备工作
		nodes = tmpNodes
		tmp = []int{}
		tmpNodes = []*TreeNode{}
		level++
	}
	return res
}

//100,19
func levelOrderIII2(root *TreeNode) [][]int {
	var (
		nodes []*TreeNode //本次遍历的nodes
		res   [][]int     //总的返回值
		tmp   []int       //缓存的本层值
		level = 1         //层次
	)
	if root != nil {
		nodes = append(nodes, root)
	}
	for len(nodes) > 0 {
		for i := len(nodes); i > 0; i-- {
			node := nodes[0]
			nodes = nodes[1:]
			if level%2 == 1 { //奇数层正常顺序放到nodes里
				tmp = append(tmp, node.Val)
			} else { //偶数层就反过来放
				tmp = append([]int{node.Val}, tmp...)
			}
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		res = append(res, tmp)
		//进入下一层的准备工作
		tmp = []int{}
		level++
	}
	return res
}
