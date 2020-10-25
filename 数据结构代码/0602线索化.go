package main

import "fmt"

const Link=0
const Thread=1

type BiThrTree struct{
	data interface{}
	lChild *BiThrTree
	lTag int
	rChild *BiThrTree
	rTag int
}

var pre *BiThrTree
//中序遍历线索化
func (bt *BiThrTree) InThreading() {
	if bt!=nil{
		//递归左子树线索化
		bt.InThreading()
		//没有左子树
		if bt.lChild==nil{
			bt.lTag=Thread
			//左子树指向前驱
			bt.lChild=pre
		}
		if pre.rChild==nil{
			pre.rTag=Thread
			//前序右子树指向后继节点（当前节点）
			pre.rChild=bt
		}
		pre=bt
		bt.InThreading()
	}
}


func (bt *BiThrTree)InOrderTraverse_Thr() {
	//p是头结点的左子树，即原树的根节点.右节点是中序遍历的最后一个节点
	p:=bt.lChild

	for p!=bt{
		//先从根节点遍历到中序遍历的第一个叶子节点
		for p.lTag==0{
			p=p.lChild
		}
		fmt.Println(p.data)
		//遍历右子树
		for p.rTag==Thread&&p.rChild!=bt{
			//p此时为某节点的子树，现在指向后驱节点，即某节点
			p=p.rChild
			fmt.Println(p.data)
		}
		p=p.rChild
	}
}