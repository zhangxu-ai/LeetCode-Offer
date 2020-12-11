package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//根据数组构造二叉树，-100表示null
func NewNodeTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	ret := &TreeNode{nums[0], nil, nil}
	tmp := []*TreeNode{ret}
	for k, v := range nums {
		if k == 0 {
			continue
		}
		if v != -100 {
			index := (k - 1) / 2
			current := &TreeNode{v, nil, nil}
			if tmp[index] != nil {
				if k%2 == 0 {
					tmp[index].Right = current
				} else {
					tmp[index].Left = current
				}
			}
			tmp = append(tmp, current)
		} else {
			tmp = append(tmp, nil)
		}
	}
	return ret
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	h, t := posOrder(root)
	if h != nil {
		h.Left = t
		t.Right = h
	}
	return h
}

//后序遍历法
func posOrder(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	//先返回左子树
	head, t1 := posOrder(root.Left)
	//后返回右子树
	h2, tail := posOrder(root.Right)
	//将root加到左子树和右子树中间
	//head和t1应该有相同的非空关系，要么全空，要么全不空
	if head == nil {
		head = root
	} else {
		t1.Right = root
	}
	if tail == nil {
		tail = root
	} else {
		h2.Left = root
	}
	root.Left = t1
	root.Right = h2
	//返回新的head和tail
	return head, tail
}

//使用全局变量，递归函数
var head, pre *TreeNode

func treeToDoublyList1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	inorder(root)
	// pre 最后变成了尾节点，需要指向头节点
	pre.Right = head
	head.Left = pre

	return head
}

//中序遍历
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	//左子树
	inorder(root.Left)
	// 初始条件下 head和pre为root
	if head == nil {
		head = root
		pre = root
	} else {
		// 前驱节点指向当前节点
		pre.Right = root
		// 当前节点指向前驱节点
		root.Left = pre
		// 前驱节点前进
		pre = root
	}

	inorder(root.Right)

}

func main() {
	a := []int{4, 2, 6, 1, 3, 5, 7}
	n := NewNodeTree(a)
	t := treeToDoublyList(n)
	for i := 0; i < len(a); i++ {
		fmt.Println(t.Val)
		t = t.Right
	}
}
