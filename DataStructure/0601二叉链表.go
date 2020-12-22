package stl

import "fmt"

type BitNode struct {
	data   interface{}
	lChild *BitNode
	rChild *BitNode
}

//前序遍历
//递归法
//考察到一个节点后，即刻输出该节点的值，并继续遍历其左右子树。(根左右)
func (bn *BitNode) PreOrderTraverse() {
	if bn == nil {
		return
	}
	fmt.Println(bn.data)
	//递归调用
	bn.lChild.PreOrderTraverse()
	bn.rChild.PreOrderTraverse()
}

//循环法
func (bn *BitNode) PreOrderTraverse1() {
	stack := NewLinkStack() //栈
	bt := bn                //游标节点
	for bt != nil || !stack.IsEmpty() {
		//一直往左走，直到左边为空
		for bt != nil {
			fmt.Println(bt)
			stack.Push(bt)
			bt = bt.lChild
		}
		//若某个节点的左子树为空，则该节点出栈，将该节点的右子树给游标节点，从而往右走
		//若栈为空，则在已经满足上面bt==nil,达到了退出大循环的条件
		//则退出
		if !stack.IsEmpty() {
			pt, _ := stack.Pop()
			bt = pt.(*BitNode).rChild
		}
	}
}

//中序排列
//考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)
//递归法
func (bn *BitNode) InOrderTraverse() {
	if bn == nil {
		return
	}
	//递归，对每个遍历到的节点都执行一次此操作，
	//先遍历左子树
	bn.lChild.InOrderTraverse()
	//到左子树为空时再访问该节点
	fmt.Println(bn.data)
	//接着遍历右子树
	bn.rChild.InOrderTraverse()
}

//循环法
//与前序不同，只有当该节点的左子树遍历完之后才访问该节点的值
func (bn *BitNode) InOrderTraverse1() {
	stack := NewLinkStack()
	bt := bn
	for bt != nil || !stack.IsEmpty() {
		for bt != nil {
			stack.Push(bt)
			bt = bt.lChild
		}
		if !stack.IsEmpty() {
			e, _ := stack.Pop()
			fmt.Println(e.(*BitNode).data)
			bt = e.(*BitNode).rChild
		}
	}
}

//后序遍历
//考察到一个节点后，将其暂存，遍历完左右子树后，再输出该节点的值。(左右根)
//递归法
func (bn *BitNode) PostOrderTraverse() {
	if bn == nil {
		return
	}
	//递归，对每个遍历到的节点都执行一次此操作，
	//先遍历左子树
	bn.lChild.InOrderTraverse()
	//接着遍历右子树
	bn.rChild.InOrderTraverse()
	//到左子树和右子树都遍历完时再访问该节点
	fmt.Println(bn.data)
}

//循环法
func (bn *BitNode) PostOrderTraverse1() {
	stack := NewLinkStack()
	bt := bn
	lastVisit := bn
	for bt != nil || !stack.IsEmpty() {
		//左边为空则开始判断右边
		for bt != nil {
			stack.Push(bt)
			bt = bt.lChild
		}
		if !stack.IsEmpty() {
			//先查看栈顶元素（当前节点元素）
			e := stack.top.data.(*BitNode)
			//如果此节点的右子树为空或者已经遍历过
			if e.rChild == nil ||
				e.rChild == lastVisit {
				fmt.Println(e.data) //访问此节点
				_, _ = stack.Pop()  //把这个节点从栈内弹出
				lastVisit = e       //设置lastVisit为当前节点
				//当前节点结束，将游标置为空，以便下一轮时，直接判断访问新的栈顶，即上一个节点元素，
				bt = nil
			} else {
				//不为空或者没有遍历过，则继续往右子树遍历
				bt = e.rChild
			}
		}
	}
}
