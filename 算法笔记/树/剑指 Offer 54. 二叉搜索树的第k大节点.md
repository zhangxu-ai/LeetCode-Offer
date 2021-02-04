# 剑指 Offer 54. 二叉搜索树的第k大节点

## 题目

给定一棵二叉搜索树，请找出其中第k大的节点。

示例 1:
输入: root = [3,1,4,null,2], k = 1

```others
   3
  / \
 1   4
  \
   2
```

输出: 4

## 思考

1. 中序遍历，生成递增数组；然后返回数组 len-k 处就行；
2. 中序遍历的倒序遍历，生成递减数组，返回数组 k-1 处

## 代码

```go
var nodes []int
var kg int

func kthLargest(root *TreeNode, k int) int {
	nodes = make([]int, 0, k)
	kg = k
	helper(root)
	return nodes[k-1]
}

//从右边开始的中序遍历
func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Right)
	nodes = append(nodes, root.Val)
	if len(nodes) >= kg {
		return
	}
	helper(root.Left)
}

```
