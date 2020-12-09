package main

import (
	"errors"
	"fmt"
)

//自己实现的栈
type SqStack09 struct {
	data []int
	top  int
}

//新建栈
func NewStack09(len int) *SqStack09 {
	return &SqStack09{
		data: make([]int, len),
		top:  -1,
	}
}

//入栈
func (s *SqStack09) Push(e int) error {
	if s.top >= len(s.data)-1 {
		return errors.New("no free space")
	}
	s.data[s.top+1] = e
	s.top++
	return nil
}

func (s *SqStack09) Pop() int {
	if s.top == -1 {
		return -1
	}
	e := s.data[s.top]
	s.top--
	return e
}

type CQueue09 struct {
	InStack  SqStack09
	OutStack SqStack09
}

//直接指定长度
func Constructor09() CQueue09 {
	return CQueue09{
		InStack:  *NewStack09(10000),
		OutStack: *NewStack09(10000),
	}
}

func (this *CQueue09) AppendTail(value int) {
	_ = this.InStack.Push(value)
}

func (this *CQueue09) DeleteHead() int {
	if this.OutStack.top > -1 {
		return this.OutStack.Pop()
	}
	for this.InStack.top > -1 {
		_ = this.OutStack.Push(this.InStack.Pop())
	}
	return this.OutStack.Pop()
}
func main() {
	c := Constructor09()
	c.AppendTail(3)
	fmt.Println(c.DeleteHead(), c.DeleteHead(), c.DeleteHead())
}
