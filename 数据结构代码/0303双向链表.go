package main

import (
	"errors"
	"fmt"
)

type DulNode struct {
	data interface{}
	next *DulNode
	prior *DulNode
}

func NewDulList(num int) (*DulNode,error){
	if num<1{
		return nil,errors.New("invalid num")
	}
	ns:=new(DulNode)
	//tn:=new(DulNode)
	ns.data=num //头结点存储链的长度
	ns.next=ns //头结点
	ns.prior=ns
	//
	//tn.data=1//暂时的第一个节点
	//tn.next=ns
	//tn.prior=ns
	//头插法
	for i:=0;i<num;i++{
		n:=new(DulNode)
		n.data=i+1
		//01新节点的前驱为第一个节点
		n.prior=ns
		//02新节点的后驱为第一个节点的后驱
		n.next=ns.next
		//03新节点的后驱的前驱为新节点
		ns.next.prior=n
		//04新节点的前驱的后驱为新节点
		ns.next=n
	}

	return ns,nil
}

func NewDulList2(num int) (*DulNode,error){
	if num<1{
		return nil,errors.New("invalid num")
	}
	ns:=new(DulNode)
	ns.data=num
	ns.next=ns //头结点
	ns.prior=ns
	//尾插法
	ne:=ns
	for i:=0;i<num;i++{
		n:=new(DulNode)
		n.data=i+1
		n.prior=ne
		n.next=ns
		ne=n
	}
	return ns,nil
}
//获取一个
func (n DulNode) GetElem(p int) (interface{}, error) {
	le:=n.data.(int)
	if p<1||p>le{
		return -1,errors.New("invalid position")
	}
	var tn *DulNode

	//在左半区，从头节点一直next
	if p<=le/2{
		tn=n.next
		for i:=1;i<p;i++{
			tn=tn.next
		}
	}else{
		//在右半区，从尾节点一直prior
		tn= n.prior
		for i:=le;i>p;i--{
			tn=tn.prior
		}
	}
	return tn.data,nil
}
//打印出所有
func (n DulNode) PrintAll() {
	for i:=1;i<=n.data.(int);i++{
		data,e:=n.GetElem(i)
		if e != nil {
			fmt.Println("get elem err:",e)
		}else{
			fmt.Println(i,"is " ,data)
		}
	}
}

//在某处插入一个
func (n *DulNode)InsertOne(e interface{},p int)bool{
	le:=n.data.(int)
	if p<1||p>le+1{
		fmt.Println("invalid position,",p)
		return false
	}
	tn:=n
	//先找到p前面的节点
	if p<=le/2{
		tn=n.next
		for i:=1;i<p-1;i++{ //从头部找，少找一次
			tn=tn.next
		}
	}else{
		//在右半区，从尾节点一直prior
		tn= n.prior
		for i:=le;i>p-1;i--{ //从尾部开始找到p的前一个节点，要多找一次
			tn=tn.prior
		}
	}
	//创建新节点
	nn:=&DulNode{
		data: e, //数据为e
		next: tn.next, //指向原先的p节点
		prior:tn,
	}
	nn.next.prior=nn
	tn.next=nn //将p前面的节点与新节点相连
	n.data=le+1 //更新节点头部的长度信息
	return true
}

//删除一个
func (n *DulNode) DeleteOne(p int) (interface{},error) {
	le:=n.data.(int)
	if p<1||p>le{
		return -1,errors.New("invalid position")
	}
	tn :=n
	//获取p的前一个节点
	if p<=le/2{
		tn=n.next
		for i:=1;i<p-1;i++{ //从头部找，少找一次
			tn=tn.next
		}
	}else{
		//在右半区，从尾节点一直prior
		tn= n.prior
		for i:=le;i>p-1;i--{ //从尾部开始找到p的前一个节点，要多找一次
			tn=tn.prior
		}
	}

	//第p个节点
	pNode:=tn.next
	//将链绕过第p个节点
	//tn.next=tn.next.next
	//p+1个节点的prior改为p-1
	pNode.next.prior=tn
	//p-1的next改为p+1
	tn.next=pNode.next
	pNode.prior=nil
	pNode.next=nil //模拟删除此节点
	n.data=n.data.(int)-1 //更新节点长度信息
	return pNode.data,nil
}

//整表删除
func (n *DulNode) Free() {
	//取第一个节点
	ns:=n.next
	var nt *DulNode
	for i:=0;i<n.data.(int);i++{
		//先备份第二个节点
		nt=ns.next
		//清空第一个节点
		ns.data=0
		ns.next=nil
		ns.prior=nil
		//第二个节点称为新的第一个节点
		ns=nt
	}
	n.data=0
	n.next=nil
}

func (n *DulNode) Push(e interface{}) bool {
	return n.InsertOne(e,n.data.(int)+1)
}

func main(){
	n,err:=NewDulList(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	n.PrintAll()
	if n.InsertOne(3.1,4){
		fmt.Println("insert ",3.1," to " ,4," success,now")
		n.PrintAll()
	}
	if d,e:= n.DeleteOne(4);e==nil{
		fmt.Println("delete ",4," success , return ",d," now")
	}else{
		fmt.Print("delete ",4," failed ,now is")
	}
	n.PrintAll()
	n.Push(6)
	n.PrintAll()
	n.Free()
	n.PrintAll()
}