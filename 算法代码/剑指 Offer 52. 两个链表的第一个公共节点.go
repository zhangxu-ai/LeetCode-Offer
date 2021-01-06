package main

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

//让长的先走到一样长，然后一起比较
//95,86
func getIntersectionNode(headA, headB *ListNode1) *ListNode1 {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		a++
		curA = curA.Next
	}
	curA = headA
	for curB != nil {
		b++
		curB = curB.Next
	}
	curB = headB
	if a > b {
		for a > b {
			curA = curA.Next
			a--
		}
	} else if a < b {
		for a < b {
			curB = curB.Next
			b--
		}
	}
	for curA != curB {
		if curA == nil || curB == nil {
			return nil
		}
		curA = curA.Next
		curB = curB.Next
	}
	return curA
}

//哈希表法，先存起来一个，然后遍历
//7,16
func getIntersectionNode1(headA, headB *ListNode1) *ListNode1 {
	if headA == nil || headB == nil {
		return nil
	}
	tmp := map[*ListNode1]bool{}
	curA, curB := headA, headB
	for curA != nil {
		tmp[curA] = true
		curA = curA.Next
	}
	for curB != nil {
		if _, ok := tmp[curB]; ok {
			return curB
		}
		curB = curB.Next
	}
	return nil
}

/*最优解法
数学规律法
两个链表长度分别为L1+C、L2+C，
C为公共部分的长度，
第一个人走了L1+C步后，回到第二个人起点走L2步；
第2个人走了L2+C步后，回到第一个人起点走L1步。
当两个人走的步数都为L1+L2+C时就两个家伙就相爱了
*/
//68,91
func getIntersectionNode2(headA, headB *ListNode1) *ListNode1 {
	cur1, cur2 := headA, headB
	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = headB
		} else {
			cur1 = cur1.Next
		}
		if cur2 == nil {
			cur2 = headB
		} else {
			cur2 = cur2.Next
		}
	}
	return cur1
}
