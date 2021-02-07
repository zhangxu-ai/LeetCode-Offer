# 剑指 Offer 07. 重建二叉树

## 题目

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出

前序遍历 `preorder =[3,9,20,15,7]`  
中序遍历 `inorder = [9,3,15,20,7]`
返回如下的二叉树：

```
 3
/ \
9  20
/  \
15   7
```

## 思考

根据前序遍历和中序遍历的特点，可以的到:

1. 前序遍历的第一个一点是根节点
2. 中序遍历中，根节点的左半部分为左子树，右半部分为右子树
3. 从根节点开始递归获取它的左右子树
    1. 如果前序列表为空，则返回 nil
    2. 如果前序列表长度为1，则返回此节点
4. 返回root节点

## 代码

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) ==0 {
		return nil
	}
	root:=new(TreeNode)
	root.Val=preorder[0]
	if len(preorder)==1{
		return root
	}
	index:=0
	for i := 0; i < len(inorder); i++ {
		if inorder[i]==preorder[0]{
			index=i
			break
		}
	}
	root.Left=buildTree(preorder[1:1+index],inorder[:index])
	root.Right=buildTree(preorder[1+index:],inorder[index+1:])
	return root
}

```
