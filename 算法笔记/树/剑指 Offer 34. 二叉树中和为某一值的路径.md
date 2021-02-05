# 剑指 Offer 34. 二叉树中和为某一值的路径

## 题目

输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

示例:
给定如下二叉树，以及目标sum = 22，

```c
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1

```

返回:

```c
[
    [5,4,11,2],
    [5,8,4,5]
]
```

## 思考

从根节点开始往下访问到叶子节点，遍历所有的路线即可。相当于前序遍历

## 代码

```go
var (
	rsp [][]int
	path []int
)

func pathSum(root *TreeNode, sum int) [][]int {
	rsp =nil
	if root==nil{
		return rsp
	}
	_pathSum(root,0,sum)
	return rsp
}

func _pathSum(root *TreeNode,now,target int) {
	now=now+root.Val
	path = append(path, root.Val)
	if root.Left!=nil{
		_pathSum(root.Left,now,target)
	}
	if root.Right!=nil{
		//pathRight:=make([]int,len(path))
		//copy(pathRight, path)
		_pathSum(root.Right,now,target)
	}
	if root.Left==nil&&root.Right==nil&&now==target{
		tmp:=make([]int,len(path))
		copy(tmp,path)
		rsp = append(rsp, tmp)
	}
	path = path[:len(path)-1]
}

```