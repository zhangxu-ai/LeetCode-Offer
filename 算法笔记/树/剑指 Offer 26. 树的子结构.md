# 剑指 Offer 26. 树的子结构

## 题目

输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

```go
     3
    / \
   4   5
  / \
 1   2
```

给定的树 B：

```go
   4
  /
 1
```

返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：`A = [1,2,3], B = [3,1]`
输出：`false`

## 思考

1. 先遍历a,找到b的根节点；然后从此根节点与b的根节点开始递归比较
2. k大解法：在递归中就进行了比较，相当于对a进行前序遍历。
   1. 先比较当前a节点的值和b的当前节点值，如果相等，则继续向下递归比较a和b的左右子树；如果不相等，则开始递归a的左右子树和b的root进行比较

## 代码

```go
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if B==nil || A==nil{
        return false
    }
    return recur(A,B)||isSubStructure(A.Left,B)||isSubStructure(A.Right,B)
}

func recur(A *TreeNode, B *TreeNode) bool{
    if B==nil {
        return true
    }
    if A==nil{
        return false
    }
    if A.Val==B.Val{
        return recur(A.Left,B.Left)&&recur(A.Right,B.Right)
    }
}
```
