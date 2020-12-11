package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//自己写的hash法
//100,17
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nh := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	ns := []*Node{nh}
	ni := map[*Node]int{head: 0}
	i := 1
	cur, cn := head.Next, nh
	//复制next
	for cur != nil {
		tmp := &Node{
			Val:    cur.Val,
			Next:   nil,
			Random: nil,
		}
		cn.Next = tmp
		cn = cn.Next
		ni[cur] = i
		ns = append(ns, cn)
		i++
		cur = cur.Next
	}
	cur = head
	cn = nh
	//复制random
	for cur != nil {
		if cur.Random == nil {
			cn.Random = nil
		} else {
			cn.Random = ns[ni[cur.Random]]
		}
		cn = cn.Next
		cur = cur.Next
	}
	return nh
}

//大佬写的hash法
//100,57
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	dic := map[*Node]*Node{}
	cur := head
	for cur != nil {
		node := &Node{
			Val:    cur.Val,
			Next:   nil,
			Random: nil,
		}
		dic[cur] = node
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		dic[cur].Next = dic[cur.Next]
		dic[cur].Random = dic[cur.Random]
		cur = cur.Next
	}
	return dic[head]
}

//大佬写的,原地复制
//100,89
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	//原地复制
	for cur != nil {
		node := &Node{
			Val:    cur.Val,
			Next:   cur.Next,
			Random: cur.Random,
		}
		cur.Next = node
		cur = node.Next
	}
	cur = head
	//修复random的关系
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	nh := head.Next
	cur = head
	//修复next关系
	for cur != nil {
		t := cur.Next
		if cur.Next != nil {
			cur.Next = cur.Next.Next
		}
		if t.Next != nil {
			t.Next = t.Next.Next
		}
		cur = cur.Next
	}
	return nh
}
