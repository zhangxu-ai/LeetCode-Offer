package main

type MaxQueue struct {
	data []int
	max  []int
}

func ConstructorMQ() MaxQueue {
	return MaxQueue{
		data: []int{},
		max:  []int{},
	}
}

func (q *MaxQueue) Max_value() int {
	if len(q.max) == 0 {
		return -1
	} else {
		return q.max[0]
	}
}

func (q *MaxQueue) Push_back(value int) {
	q.data = append(q.data, value)
	if len(q.max) == 0 {
		q.max = append(q.max, value)
	} else {
		i := 0
		for i = 0; i < len(q.max) && q.max[i] >= value; i++ {
		}
		q.max[i] = value
		q.max = q.max[:i+1]
	}
}

func (q *MaxQueue) Pop_front() int {
	if len(q.data) == 0 {
		return -1
	}
	t := q.data[0]
	q.data = q.data[1:]
	if t == q.max[0] {
		q.max = q.max[1:]
	}
	return t
}
