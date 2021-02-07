package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	if len(preorder) == 1 {
		return root
	}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}
	root.Left = buildTree(preorder[1:1+index], inorder[:index])
	root.Right = buildTree(preorder[1+index:], inorder[index+1:])
	return root
}
