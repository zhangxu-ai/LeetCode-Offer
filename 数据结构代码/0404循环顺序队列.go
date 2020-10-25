package main

import (
	"errors"
	"fmt"
)

const MaxSize2 = 999
type SqQueue struct{
	queue [MaxSize2]interface{}
	front int
	rear int
}

func NewSqQueue() *SqQueue{
	return &SqQueue{
		queue: [999]interface{}{},
		front: 0,
		rear:  0,
	}
}
func (sq *SqQueue) Length() int {
	return (sq.rear-sq.front+MaxSize2)%MaxSize2
}

func (sq *SqQueue) EnQueue(e interface{}) error {
	if (sq.rear+1)%MaxSize2==sq.front{
		return errors.New("no free space")
	}
	sq.queue[sq.rear]=e
	//sq.rear++
	//if sq.rear>=MaxSize2{
	//	sq.rear=sq.rear-MaxSize2
	//}
	//real后移，到尾部则从头开始
	sq.rear=(sq.rear+1)%MaxSize2
	return nil
}

func (sq *SqQueue) DeQueue() (interface{},error) {
	if sq.Length()==0{
		return nil, errors.New("no element now")
	}
	e:=sq.queue[sq.front]
	sq.queue[sq.front]=0
	//front后移，到尾部则从头开始
	sq.front=(sq.front+1)%MaxSize2
	return e, nil
}

func main(){
	sq:=NewSqQueue()
	for i:=0;i<1000;i++{
		if err:=sq.EnQueue(i);err!=nil{
			fmt.Println(i,err)
		}
	}
	for i:=0;i<500;i++{
		_,_=sq.DeQueue()
		//fmt.Println(e,err)
	}
	for i:=998;i<1098;i++{
		if err:=sq.EnQueue(i);err!=nil{
			fmt.Println(i,err)
		}
	}

	fmt.Println(*sq)
}