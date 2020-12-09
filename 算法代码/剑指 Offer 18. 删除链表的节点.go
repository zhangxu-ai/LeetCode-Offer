package main

type ListNodeO struct {
	Val  int
	Next *ListNodeO
}

func deleteNode(head *ListNodeO, val int) *ListNodeO {
	//第一个节点就是
	if head.Val == val {
		return head.Next
	}
	//在中间部分
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			return head
		}
		cur = cur.Next
	}
	//不存在，直接返回原链表
	return head
}

//递归,占用内存较大
func deleteNode1(head *ListNodeO, val int) *ListNodeO {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode1(head.Next, val)
	return head
}
