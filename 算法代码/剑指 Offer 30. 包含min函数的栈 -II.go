package main

//69,60
type MinStack2 struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor3() MinStack2 {
	return MinStack2{
		data: []int{},
		min:  []int{},
	}
}

func (s *MinStack2) Push(x int) {
	s.data = append(s.data, x)
	l := len(s.min)
	if l > 0 {
		if s.min[l-1] >= x {
			s.min = append(s.min, x)
		}
	} else {
		s.min = append(s.min, x)
	}
}

func (s *MinStack2) Pop() {
	l := len(s.data)
	if l == 0 {
		return
	}
	if s.data[l-1] == s.min[len(s.min)-1] {
		s.min = s.min[:len(s.min)-1]
	}
	s.data = s.data[:l-1]
}

func (s *MinStack2) Top() int {
	l := len(s.data)
	if l == 0 {
		return 0
	}
	return s.data[l-1]
}

func (s *MinStack2) Min() int {
	l := len(s.min)
	if l == 0 {
		return 0
	}
	return s.min[l-1]
}
