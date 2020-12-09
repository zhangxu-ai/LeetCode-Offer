package main

//DFS
func isSymmetric(root *TreeNode) bool {
	return recurS(root, root)
}

func recurS(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return recurS(a.Left, b.Right) && recurS(a.Right, b.Left)
}

//BFS
func isSymmetric1(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	nodes := []*TreeNode{root}
	c, l := 0, 1 //计数、层
	for c < len(nodes) {
		if nodes[c] != nil {
			nodes = append(nodes, nodes[c].Left, nodes[c].Right)
		} else {
			nodes = append(nodes, nil, nil)
		}
		//每一层的最后一个元素
		e := int(myPowS(2, l)) - 2
		if c == e {
			//到每一层的第一个元素
			et := e
			s := int(myPowS(2, l-1)) - 1
			mid := (e - s) / 2
			nc := 0
			//从此层的第一个元素开始，到最后一个元素
			for i := s; i <= s+mid; i++ {
				if nodes[i] == nil && nodes[e] == nil {
					if i == e {
						nc++
					} else {
						nc += 2
					}
				} else if nodes[i] == nil || nodes[e] == nil || nodes[i].Val != nodes[e].Val {
					return false
				}
				e--
			}
			//如果这一层全都是nil，说明已经便利到了最后一层的后一层，可以结束循环了
			if nc == et-s+1 {
				break
			}
		} else if c > e { //到下一层
			l++
		}
		//计数器加一
		c++
	}
	return true
}

func myPowS(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	var r = 1.0
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n > 0 {
		if (n & 1) == 1 {
			r *= x
		}
		x *= x
		n >>= 1
	}
	return r
}
