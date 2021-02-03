package main

import (
	"errors"
	"fmt"
)

//准备工作
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历链表
func (bn *TreeNode) PreOrderTraverse() {
	if bn == nil {
		return
	}
	fmt.Println(bn.Val)
	//递归调用
	bn.Left.PreOrderTraverse()
	bn.Right.PreOrderTraverse()
}

//自己实现的栈
type SqStack07 struct {
	data []*TreeNode
	top  int
}

//新建栈
func NewStack07(len int) *SqStack07 {
	return &SqStack07{
		data: make([]*TreeNode, len),
		top:  -1,
	}
}

//入栈
func (s *SqStack07) Push(e *TreeNode) error {
	if s.top >= len(s.data)-1 {
		return errors.New("no free space")
	}
	s.data[s.top+1] = e
	s.top++
	return nil
}

//出栈
func (s *SqStack07) Pop() (*TreeNode, error) {
	if s.top == -1 {
		return nil, errors.New("stack no elements")
	}
	e := s.data[s.top]
	s.top--
	return e, nil
}

//重建链表1
func buildTree(preorder []int, inorder []int) *TreeNode {
	for i := range inorder {
		if inorder[i] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:i+1], inorder[:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil
}

//重建链表2
//使用栈的方法实现
/*
如果使用栈来解决首先要搞懂一个知识点，就是前序遍历挨着的两个值比如m和n，他们会有下面两种情况之一的关系。=
1，n是m左子树节点的值。
2，n是m右子树节点的值或者是m某个祖先节点的右节点的值。
	1. 对于第一个知识点我们很容易理解，如果m的左子树不为空，那么n就是m左子树节点的值。
	2. 对于第二个问题，如果一个结点没有左子树只有右子树，那么n就是m右子树节点的值，\
	如果一个结点既没有左子树也没有右子树，那么n就是m某个祖先节点的右节点，我们只要找到这个祖先节点就好办了。
*/
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	s := NewStack07(len(preorder))
	//根节点就是前序遍历的第一个元素
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	cur := root
	/*
		preorder := []int{3,9,20,15,7}
		inorder := []int{9,3,15,20,7}
	*/
	//从左到右遍历前序列表，并且根据中序遍历列表的关系，
	//确定每一个前序元素和它上一个元素的关系，从而构造合适的二叉树
	//假设第一轮循环的情况
	//i一直向右走，j只有在
	for i, j := 1, 0; i < len(preorder); i++ {
		//第一种情况
		//跟节点的值不等于中序遍历的第一个值，说明此节点有左子树；
		//如果没有左子树，则中序遍历第一个元素就是根节点
		if (cur.Val) != inorder[j] {
			//这时前序遍历中，第一个元素是根节点的值,第二个元素就是他的左子树
			cur.Left = &TreeNode{
				Val:   preorder[i],
				Left:  nil,
				Right: nil,
			}
			//先将当前节点入栈，即将3(根节点)入栈
			_ = s.Push(cur)
			//继续向左遍历，cur此时为9，即找左子树节点的左右子树
			cur = cur.Left
		} else { //第二种情况，在循环的第二轮才会触发，
			j++
			//找到这个祖先节点
			//每个入栈的元素都有左子树，即该元素在中序列表中，前面一定有元素
			//当栈不为空且栈顶元素的值等于中序的第二值3
			for s.top != -1 && s.data[s.top].Val == inorder[j] {
				//cur出栈
				cur, _ = s.Pop()
				//
				j++
			}
			cur.Right = &TreeNode{
				Val:   preorder[i],
				Left:  nil,
				Right: nil,
			}
			cur = cur.Right
		}
	}
	return root
}
func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	n := buildTree1(preorder, inorder)
	n.PreOrderTraverse()
}
