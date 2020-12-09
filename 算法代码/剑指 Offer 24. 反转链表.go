package main

type ListNodeR struct {
	Val  int
	Next *ListNodeR
}

//双指针法 迭代
func reverseList(head *ListNodeR) *ListNodeR {
	var newHead, tmp *ListNodeR
	for head != nil {
		tmp = head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}
	return newHead
}

//递归
func reverseList1(head *ListNodeR) *ListNodeR {
	if head == nil || head.Next == nil {
		return head
	}
	r := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return r
}
