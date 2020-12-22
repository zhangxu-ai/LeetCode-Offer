package stl

import (
	"errors"
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

func NewList(num int) (*node, error) {
	if num < 1 {
		return nil, errors.New("invalid num")
	}
	ns := new(node)
	ns.data = num //头结点存储链的长度
	ns.next = nil //头结点
	//头插法
	for i := 0; i < num; i++ {
		n := new(node)
		n.data = i + 1
		n.next = ns.next
		ns.next = n
	}
	return ns, nil
}

func NewList2(num int) (*node, error) {
	if num < 1 {
		return nil, errors.New("invalid num")
	}
	ns := new(node)
	ns.data = num
	ns.next = nil //头结点
	ne := ns      //尾节点
	//尾插法
	for i := 0; i < num; i++ {
		n := new(node)
		n.data = i + 1
		ne.next = n
		ne = n
	}
	ne.next = nil
	return ns, nil
}

//获取一个
func (n node) GetElem(p int) (interface{}, error) {
	if p < 1 || p > n.data.(int) {
		return -1, errors.New("invalid position")
	}
	tn := n.next
	for i := 1; i < p; i++ {
		tn = tn.next
	}
	return tn.data, nil
}

//打印出所有
func (n node) PrintAll() {
	for i := 1; i <= n.data.(int); i++ {
		data, e := n.GetElem(i)
		if e != nil {
			fmt.Println("get elem err:", e)
		} else {
			fmt.Println(i, "is ", data)
		}
	}
}

//在某处插入一个
func (n *node) InsertOne(e interface{}, p int) bool {
	if p < 1 || p > n.data.(int)+1 {
		fmt.Println("invalid position,", p)
		return false
	}
	tn := n
	//先找到p前面的节点
	for i := 1; i < p; i++ {
		tn = tn.next
	}
	//创建新节点
	nn := &node{
		data: e,       //数据为e
		next: tn.next, //指向原先的p节点
	}
	tn.next = nn              //将p前面的节点与新节点相连
	n.data = n.data.(int) + 1 //更新节点头部的长度信息
	return true
}

//删除一个
func (n *node) DeleteOne(p int) (interface{}, error) {
	if p < 1 || p > n.data.(int) {
		return -1, errors.New("invalid position")
	}
	tn := n
	//获取p的前一个节点
	for i := 1; i < p; i++ {
		tn = tn.next
	}
	//第p个节点
	pNode := tn.next
	//将链绕过第p个节点
	//tn.next=tn.next.next
	tn.next = pNode.next
	pNode.next = nil          //模拟删除此节点
	n.data = n.data.(int) - 1 //更新节点长度信息
	return pNode.data, nil
}

//整表删除
func (n *node) Free() {
	ns := n.next
	var nt *node
	for i := 0; i < n.data.(int); i++ {
		nt = ns.next
		ns.data = 0
		ns.next = nil
		ns = nt
	}
	n.data = 0
	n.next = nil
}

func (n *node) Push(e interface{}) bool {
	return n.InsertOne(e, n.data.(int)+1)
}

func main() {
	n, err := NewList2(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	n.PrintAll()
	if n.InsertOne(3.1, 4) {
		fmt.Println("insert ", 3.1, " to ", 4, " success,now")
		n.PrintAll()
	}
	if d, e := n.DeleteOne(4); e == nil {
		fmt.Println("delete ", 4, " success , return ", d, " now")
	} else {
		fmt.Print("delete ", 4, " failed ,now is")
	}
	n.PrintAll()
	n.Push(6)
	n.PrintAll()
	n.Free()
	n.PrintAll()
}
