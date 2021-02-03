package main

import (
	"fmt"
)

//自己实现的栈
type SqStack091 struct {
	data []int
	top  int
}

//新建栈
func NewStack(len int) *SqStack091 {
	return &SqStack091{
		data: make([]int, len),
		top:  -1,
	}
}

//入栈
func (s *SqStack091) Push(e int) error {
	//if s.top >= len(s.data)-1 {
	//	return errors.New("no free space")
	//}
	s.data = append(s.data, e)
	s.top++
	return nil
}

//出栈
func (s *SqStack091) Pop() int {
	if s.top == -1 {
		return -1
	}
	e := s.data[s.top]
	s.data = s.data[0:s.top]
	s.top--
	return e
}

type CQueue struct {
	InStack  SqStack091
	OutStack SqStack091
}

//直接指定长度
func ConstructorQ() CQueue {
	return CQueue{
		InStack:  *NewStack(10000),
		OutStack: *NewStack(10000),
	}
}

//动态扩容
func Constructor1() CQueue {
	return CQueue{
		InStack: SqStack091{
			data: []int{},
			top:  -1,
		},
		OutStack: SqStack091{
			data: []int{},
			top:  -1,
		},
	}
}

func (this *CQueue) AppendTail(value int) {
	_ = this.InStack.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.OutStack.top > -1 {
		return this.OutStack.Pop()
	}
	for this.InStack.top > -1 {
		_ = this.OutStack.Push(this.InStack.Pop())
	}
	return this.OutStack.Pop()
}
func main() {
	c := Constructor1()
	c.AppendTail(3)
	fmt.Println(c.DeleteHead(), c.DeleteHead(), c.DeleteHead())
}
