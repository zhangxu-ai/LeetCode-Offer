package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
