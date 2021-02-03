package main

type CQueue struct {
	a []int
	b []int
}

func Constructor() CQueue {
	return CQueue{
		a: []int{},
		b: []int{},
	}
}

func (q *CQueue) AppendTail(value int) {
	q.a = append(q.a, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.a) == 0 && len(q.b) == 0 {
		return -1
	} else if len(q.a) != 0 && len(q.b) == 0 {
		//模拟a出栈，b入栈过程
		for len(q.a) > 0 {
			q.b = append(q.b, q.a[len(q.a)-1])
			q.a = q.a[:len(q.a)-1]
		}
	}
	t := q.b[len(q.b)-1]
	q.b = q.b[:len(q.b)-1]
	return t
}
