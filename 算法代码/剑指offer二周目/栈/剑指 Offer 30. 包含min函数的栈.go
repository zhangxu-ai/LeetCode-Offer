package main

type MinStack struct {
	data []int
	min  []int
}

func ConstructorMS() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{},
	}
}

func (s *MinStack) Push(x int) {
	s.data = append(s.data, x)
	if len(s.min) == 0 || s.min[len(s.min)-1] >= x {
		s.min = append(s.min, x)
	}
}

func (s *MinStack) Pop() {
	l := len(s.data)
	if l == 0 {
		return
	}
	if s.data[l-1] == s.min[len(s.min)-1] {
		s.min = s.min[:len(s.min)-1]
	}
	s.data = s.data[:l-1]
}

func (s *MinStack) Top() int {
	l := len(s.data)
	if l == 0 {
		return -1
	}
	return s.data[l-1]
}

func (s *MinStack) Min() int {
	l := len(s.min)
	if l == 0 {
		return -1
	}
	return s.min[l-1]
}
