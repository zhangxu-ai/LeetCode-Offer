package main

/*
尴尬，以为上一题是最后一题，都撒花了
*/
//后续遍历此树，如果
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	//越过叶子节点
	if root == nil {
		return nil
	}
	//当前节点就是p或q，则可能是公共祖先
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//如果p,q分别在root的左或右节点，则
	if left != nil && right != nil {
		return root
	}
	//如果只有一个不存在，则继续往上
	if left == nil {
		return right
	}
	return left
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}

//2021.01.25 19:46 于盐城一刷完毕
