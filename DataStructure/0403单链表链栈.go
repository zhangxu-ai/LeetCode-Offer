package stl

import (
	"errors"
	"fmt"
)

type StackNode struct {
	data interface{}
	next *StackNode
}
type LinkStack struct {
	top   *StackNode
	count int
}

func NewLinkStack() *LinkStack {
	return &LinkStack{
		top:   nil,
		count: 0,
	}
}

func (ls *LinkStack) Push(e interface{}) {
	sn := &StackNode{
		data: e,
		next: ls.top, //指向旧的头部
	}
	ls.top = sn //sn成为新的头部
	ls.count++
	return
}

func (ls *LinkStack) Pop() (interface{}, error) {
	if ls.count == 0 || ls.top == nil {
		return nil, errors.New("stack no elements")
	}
	e := ls.top.data
	ls.top = ls.top.next
	ls.count--
	return e, nil
}

func (ls LinkStack) IsEmpty() bool {
	return ls.count == 0 && ls.top == nil
}

func main() {
	ls := NewLinkStack()
	ls.Push(1)
	fmt.Println("push 1 success")
	if e, err := ls.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(e)
	}
}
