package stl

import "errors"

const MaxSize = 999

type SqStack struct {
	data [MaxSize]interface{}
	top  int
}

func (s *SqStack) Push(e interface{}) error {
	if s.top >= len(s.data)-1 {
		return errors.New("no free space")
	}
	s.data[s.top+1] = e
	s.top++
	return nil
}

func (s *SqStack) Pop() (interface{}, error) {
	if s.top == -1 {
		return -1, errors.New("stack no elements")
	}
	e := s.data[s.top]
	s.top--
	return e, nil
}
