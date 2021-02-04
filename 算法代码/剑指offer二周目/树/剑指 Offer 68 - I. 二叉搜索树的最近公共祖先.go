package main

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for {
		if cur.val > p.val && cur.val > q.val {
			cur = cur.left
		} else if cur.val < p.val && cur.val < q.val {
			cur = cur.right
		} else {
			return cur
		}
	}

}
