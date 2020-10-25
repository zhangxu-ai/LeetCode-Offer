package main

import (
	"errors"
	"fmt"
)

type QNode struct{
	data interface{}
	next *QNode
}
type LinkQueue struct{
	font *QNode
	rear *QNode
}

func NewLinkQueue() *LinkQueue{
	top:=&QNode{
		data: nil,
		next: nil,
	}
	return &LinkQueue{
		font: top,
		rear: top,
	}
}

func (lq *LinkQueue) EnQueue(e interface{}) error {
	n:=&QNode{
		data: e,
		next: nil,
	}
	//尾部指向新节点
	lq.rear.next=n
	//更新尾部
	lq.rear=n
	return nil
}

func (lq *LinkQueue) DeQueue() (interface{},error) {
	if lq.font.next==nil{
		return nil, errors.New("no element now")
	}
	e:=lq.font.next.data
	//font指向原来的第二个节点，即为删除了第一个节点
	lq.font.next=lq.font.next.next
	//如果只有一个元素，删除后将rear节点指回font节点
	if lq.font.next==nil{
		lq.rear=lq.font
	}
	return e, nil
}

func main(){
	lq:=NewLinkQueue()
	_=lq.EnQueue(9)
	e,err:=lq.DeQueue()
	fmt.Println(e,err)
}