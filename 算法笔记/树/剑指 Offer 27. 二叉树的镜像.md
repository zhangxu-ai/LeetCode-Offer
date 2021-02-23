# 剑指 Offer 27. 二叉树的镜像

## 题目

请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

```go
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```

镜像输出：

```go
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

## 思考

1. 后续遍历即可。从叶子节点开始互换左右子树，然后一层层向上互换

## 代码

```go
func mirrorTree(root *TreeNode) *TreeNode {
    if root!=nil{
    root.Left,root.Right=mirrorTree(root.Right),mirrorTree(root.Left)
    }
    return root
}
```
