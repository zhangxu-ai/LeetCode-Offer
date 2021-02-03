package main

type ListNodeK struct {
	Val  int
	Next *ListNodeK
}

func getKthFromEnd(head *ListNodeK, k int) *ListNodeK {
	c, cur := 1, head.Next
	for cur != nil {
		c++
		cur = cur.Next
	}
	cur = head
	for i := 0; i < c-k; i++ {
		cur = cur.Next
	}
	return cur
}

func getKthFromEnd1(head *ListNodeK, k int) *ListNodeK {
	var tmp []*ListNodeK
	for head != nil {
		tmp = append(tmp, head)
		head = head.Next
	}
	return tmp[len(tmp)-k] //解题时不需要考虑越界情况
}

//反转链表不对，但是可以复习以下反转链表的写法
//反转链表
//func getKthFromEnd2(head *ListNodeK, k int) *ListNodeK {
//	var newHead *ListNodeK
//	for head!=nil{
//		//缓存第二个节点
//		tmp:=head.Next
//		//将第一个节点使用头插法插入到新的链表的头部
//		head.Next=newHead
//		newHead=head
//		//恢复旧链表的头部
//		head=tmp
//	}
//	for i:=0;i<k;i++ {
//		newHead=newHead.Next
//	}
//	return newHead
//}

//双指针解法
func getKthFromEnd3(head *ListNodeK, k int) *ListNodeK {
	f, s := head, head
	//s先向前跑k步
	for i := 0; i < k; i++ {
		s = s.Next
	}
	//同时向前跑，直到s跑到头，这时f就是倒数第k个
	for s != nil {
		s = s.Next
		f = f.Next
	}
	return f
}
