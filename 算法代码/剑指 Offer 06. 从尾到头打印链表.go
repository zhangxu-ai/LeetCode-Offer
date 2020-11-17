package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var a []int
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}
	s := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		s[len(a)-1-i] = a[i]
	}
	return s
}

//递归，虽然看起来很但是时间和空间复杂度太高了，😂😂😂
func reversePrint1(head *ListNode) []int {
	var s []int
	if head != nil {
		s1 := reversePrint1(head.Next)
		s = append(s, s1...)
		s = append(s, head.Val)
	}
	return s
}

//这种方式看似麻烦，但是时间和空间都是100%
//可能因为go的内部切片如果不指定大小，依靠动态扩容的话，内存和性能损耗会比较多
//第一次循环看似时间很多，但是是指针赋值操作，很快，
//确定了个数，再去创建固定长度的切片，可以直接通过index访问，不需要append老追加，加快了速度
func reversePrint2(head *ListNode) []int {
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	cur = head
	s := make([]int, count)
	for i := 0; cur != nil; i++ {
		s[count-i-1] = cur.Val
		cur = cur.Next
	}
	return s
}

//先顺序遍历然后反转数组
func reversePrint3(head *ListNode) []int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return arr
}

//反转链表，然后顺序遍历
func reversePrint4(head *ListNode) []int {
	var (
		newHead *ListNode
		count   int
		s       []int
	)
	//定义一个结点node作为"临时中转站"，初始化与否并无大影响。
	/*
		在当前指针不为NULL时，先对原链表做头删操作，再对新链表做头插操作。
	*/
	for head != nil {
		//先换成第二个节点
		tmp := head.Next
		//将head的next指向新的节点
		head.Next = newHead
		//新节点指向旧节点
		newHead = head
		//旧head更新到第二节点
		head = tmp
		count++
	}
	for newHead != nil {
		s = append(s, newHead.Val)
		newHead = newHead.Next
	}
	return s
}
func main() {

}
