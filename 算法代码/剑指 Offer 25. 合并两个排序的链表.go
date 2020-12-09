package main

//较慢，内存少
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	//以l1为基准,将l2的元素不断向l1里面插入
	cur := l1
	l2t := l2
	for l2 != nil {
		if cur.Val <= l2.Val {
			if cur.Next == nil || cur.Next.Val >= l2.Val {
				//缓存l2的下一个头
				l2t = l2.Next
				//插入
				l2.Next = cur.Next
				cur.Next = l2
				//都向前移动
				l2 = l2t
			}
			cur = cur.Next
		} else {
			//只有在l1的头部大于l2的头部时才会才会出现这种情况
			//只会出现一次这种情况
			//缓存l2的下一个头
			l2t = l2.Next
			//在cur前面加一个node
			l2.Next = cur
			cur = l2
			//注意要更新l1
			l1 = cur
			//l2向后一步
			l2 = l2t
		}
	}
	return l1
}

//创建一个新的链表，并返回
//不知道为啥内存使用比上一个少
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	towInOne := new(ListNode)
	cur := towInOne
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			//更新cur
			cur.Next = l1
			cur = cur.Next
			//l1后移
			l1 = l1.Next
		} else {
			//更新cur
			cur.Next = l2
			cur = cur.Next
			//l2后移
			l2 = l2.Next
		}
	}

	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return towInOne.Next
}

//递归用法
//快，内存多
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	head := new(ListNode)
	if l1.Val <= l2.Val {
		head = l1
		head.Next = mergeTwoLists2(l1.Next, l2)
	} else {
		head = l2
		head.Next = mergeTwoLists2(l1, l2.Next)
	}
	return head
}
