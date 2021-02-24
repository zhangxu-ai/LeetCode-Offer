# 剑指 Offer 32 - II. 从上到下打印二叉树 II

## 题目

从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

例如:
给定二叉树: [3,9,20,null,null,15,7],

```go
    3
   / \
  9  20
    /  \
   15   7
```

返回其层次遍历结果：

```c
[
  [3],
  [9,20],
  [15,7]
]
```

## 思考

这一题就是之前做的树里面的一题，在[序列号二叉树](../树/剑指%20Offer%2037.%20序列化二叉树.md)中我们已经做过了。不过这一次在队列上需要变换一下，在每一层缓存下一层的节点时，要存到另一个队列中。这次再复习一下广度优先遍历

## 代码

```go
func levelOrder(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	var res [][]int
	var tmp []int
	var queue = []*TreeNode{root}
	var next []*TreeNode
	for len(queue)!=0 {
		if queue[0]!=nil{
			tmp = append(tmp, queue[0].Val)
			next = append(next, queue[0].Left,queue[0].Right)
		}
		queue=queue[1:]
		if len(queue)==0&&len(next)!=0{
			queue,next=next,queue
			res = append(res, tmp)
			tmp=[]int{}
		}
	}
	return res
}

```
