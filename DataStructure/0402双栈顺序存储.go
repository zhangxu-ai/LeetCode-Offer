package stl

import (
	"errors"
	"fmt"
	"strconv"
)

const MaxSize1 = 999

type SqDulStack struct {
	data [MaxSize1]interface{}
	top1 int
	top2 int
}

func NewSqDulStack() *SqDulStack {
	return &SqDulStack{
		data: [999]interface{}{},
		top1: 0,
		top2: MaxSize1 - 1,
	}
}

func (s *SqDulStack) Push(e interface{}, stackNum int) error {
	if s.top1+1 == s.top2 {
		return errors.New("no free space")
	}
	if stackNum == 1 {
		s.data[s.top1+1] = e
		s.top1++
	} else if stackNum == 2 {
		s.data[s.top2-1] = e
		s.top2--
	} else {
		return errors.New("invalid stackNum: only be 1 or 2")
	}
	return nil
}

func (s *SqDulStack) Pop(stackNum int) (interface{}, error) {
	var e interface{}
	if stackNum == 1 {
		if s.top1 == -1 {
			return -1, errors.New("stack" + strconv.Itoa(stackNum) + " no elements")
		}
		e = s.data[s.top1]
		s.top1--
	} else if stackNum == 2 {
		if s.top2 == MaxSize1-1 {
			return -1, errors.New("stack" + strconv.Itoa(stackNum) + " no elements")
		}
		e = s.data[s.top2]
		s.top2++
	} else {
		return -1, errors.New("invalid stackNum: only be 1 or 2")
	}
	return e, nil
}

func main() {
	s := NewSqDulStack()
	err := s.Push(1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("push 1 to 1 success")
	}

	err1 := s.Push(1, 2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("push 1 to 2 success")
	}

	e, err3 := s.Pop(1)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(e)
	}

	e2, err4 := s.Pop(2)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(e2)
	}

	fmt.Println(s.top1, s.top2)
}
