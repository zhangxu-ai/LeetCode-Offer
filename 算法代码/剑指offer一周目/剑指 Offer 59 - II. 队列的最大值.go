package main

import "fmt"

//和30差不多
type MaxQueue struct {
	queue []int
	max   []int
}

func ConstructorMq() MaxQueue {
	return MaxQueue{
		queue: []int{},
		max:   []int{},
	}
}

func (q *MaxQueue) Max_value() int {
	l := len(q.max)
	if l == 0 {
		return -1
	}
	return q.max[0]
}

func (q *MaxQueue) Push_back(value int) {
	q.queue = append(q.queue, value)
	l := len(q.max)
	i := l - 1
	for i >= 0 && q.max[i] < value {
		i--
	}
	if l == 0 || i == l-1 {
		q.max = append(q.max, value)
		return
	}
	q.max = q.max[:i+1]
	q.max = append(q.max, value)
}

func (q *MaxQueue) Pop_front() int {
	if len(q.queue) == 0 {
		return -1
	}
	r := q.queue[0]
	q.queue = q.queue[1:]
	if q.max[0] == r {
		q.max = q.max[1:]
	}
	return r
}

func main() {
	mq := ConstructorMq()
	mq.Push_back(1)
	fmt.Println(mq.max)

	mq.Push_back(2)
	fmt.Println(mq.max)

	mq.Push_back(4)
	fmt.Println(mq.max)

	mq.Push_back(3)
	fmt.Println(mq.max)
	mq.Push_back(5)
	fmt.Println(mq.max)

	mq.Pop_front()
	fmt.Println(mq.max)
	mq.Pop_front()
	fmt.Println(mq.max)
	mq.Pop_front()
	fmt.Println(mq.queue)
	fmt.Println(mq.max)

}
