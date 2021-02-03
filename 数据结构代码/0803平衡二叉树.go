package stl

type AvlTree struct {
	data   int
	lChild *AvlTree
	rChild *AvlTree
	//height int //深度
}

//求节点高度、递归
func (at *AvlTree) Height() int {
	if at == nil {
		return 0
	}
	rh := at.rChild.Height()
	lh := at.lChild.Height()
	if rh >= lh {
		return rh + 1
	} else {
		return lh + 1
	}
}

//左旋，返回新的根节点
func (at *AvlTree) LeftR() *AvlTree {
	tt := at.rChild
	at.rChild = tt.lChild
	tt.lChild = at
	//更新高度
	//at.height=at.Height()
	//tt.height=at.Height()
	return tt
}

//右旋，返回新的根节点
func (at *AvlTree) RightR() *AvlTree {
	tt := at.lChild
	at.lChild = tt.rChild
	tt.rChild = at
	//at.height=at.Height()
	//tt.height=tt.Height()
	return tt
}

func (at *AvlTree) GetBalance() int {
	if at == nil {
		return 0
	}
	return at.lChild.Height() - at.rChild.Height()
}

func (at *AvlTree) InsertOne(data int) *AvlTree {

	//先顺序插入数据
	if at == nil {
		return &AvlTree{
			data:   data,
			lChild: nil,
			rChild: nil,
			//height: 1,
		}
	}
	if at.data == data {
		return at
	} else if data < at.data {
		at.lChild = at.lChild.InsertOne(data)
	} else {
		at.rChild = at.rChild.InsertOne(data)
	}
	//更新节点高度
	//at.height=at.Height()
	//获取balance平衡值
	b := at.GetBalance()
	//LL情况：左边高，且值插入到了左子树的左边
	if b > 1 {
		if data < at.lChild.data {
			return at.RightR()
		} else if data > at.lChild.data {
			//LR情况，先把该节点的左子树向左旋
			at.lChild = at.lChild.LeftR()
			return at.RightR()
		}
	} else if b < -1 {
		//RR的情况，当前节点左旋
		if data > at.rChild.data {
			return at.LeftR()
		} else if data > at.lChild.data {
			//RL情况，先把该节点的右子树向右旋，该节点再左旋
			at.lChild = at.rChild.RightR()
			return at.LeftR()
		}
	}
	return at
}

//判断是否平衡
func (at *AvlTree) IsBalanced() bool {
	if at == nil {
		return true
	}
	lh := at.lChild.Height()
	rh := at.rChild.Height()
	//判断左右子树高度之差是否小于1，并且结点的左右子树平衡，返回true;
	if ((lh-rh) < 1 || (lh-rh) > -1) && (at.lChild.IsBalanced() && at.rChild.IsBalanced()) {
		return true
	} else {
		return false
	}
}

//对上述算法的改进
//当输入是一颗斜树的时候，其时间复杂度将变成O(n^2) 。
//问题在于我们判断二叉树是否平衡的函数 isBalanced() 当中嵌套了一个计算树的高度的函数height() ，
//这样一来，当树为一颗斜树的时候，时间复杂度就会达到O(n^2)  。
//解决的办法就是将这两个函数合并，取消单独调用的height（）函数，而是在递归进行判断的时候计算树的高度。
func (at *AvlTree) IsBalancedUtil(h *int) bool {
	if at == nil {
		*h = 0
		return true
	}
	//左右子树的高度
	var lh, rh int
	//左右子树是否平衡的标志
	var l, r bool
	l = at.lChild.IsBalancedUtil(&lh)
	r = at.rChild.IsBalancedUtil(&rh)
	//根据返回的lh和rh，计算自身高度
	if lh >= rh {
		*h = lh + 1
	} else {
		*h = rh + 1
	}
	if !(l && r) {
		return false
	}
	if (lh-rh) < -1 || (lh-rh) > 1 {
		return false
	}
	return true
}
func DeleteOne(at *AvlTree, data int) *AvlTree {
	//第一步：标准的删除BST操作
	if at == nil {
		return at
	}
	if data < at.data {
		at.lChild = DeleteOne(at.lChild, data)
	} else if data > at.data {
		at.rChild = DeleteOne(at.rChild, data)
	} else {
		//先判断只有一个子节点或者无节点的情况
		if at.rChild == nil || at.lChild == nil {
			if at.rChild != nil {
				*at = *at.rChild
			} else if at.lChild != nil {
				*at = *at.lChild
			} else {
				//无子节点，删除当前节点
				at = nil
			}
		} else {
			//既有左子树又有右子树,找到其直接后继节点，
			//即右子树的最左节点
			q := at
			t := at.rChild
			for t.lChild != nil {
				q = t //慢指针，始终指向t的双亲节点
				t = t.lChild
			}
			//将后继节点的值赋给此节点
			at.data = t.data
			//说明没有左走。即该节点的右子树就是他的直接后继节点，
			//也说明右子树没有左节点,直接将该节点的右子树更新为直接后继节点的右子树，替换掉后继节点
			if q == at {
				q.rChild = t.rChild
			} else {
				//向左走了，直到后继节点没有左子树，说明后继节点只有右子树或者没有节点，
				//将后继节点的右子树接到后继节点的左子树上，替换掉后继节点
				q.lChild = t.rChild
			}
		}
	}
	if at == nil {
		return at
	}
	//第二步，获取删除节点的平衡因子
	b := at.GetBalance()
	//不平衡的4种情况
	//因为在递归函数之后，从插入节点那一层之上所有的层都会执行检验平衡度的操作
	//LL
	if b > 1 && at.lChild.GetBalance() >= 0 {
		return at.RightR()
	} else if b > 1 && at.lChild.GetBalance() < 0 {
		at.lChild = at.lChild.LeftR()
		return at.RightR()
	} else if b < -1 && at.rChild.GetBalance() <= 0 {
		return at.LeftR()
	} else if b < -1 && at.rChild.GetBalance() > 0 {
		at.rChild = at.rChild.RightR()
		return at.LeftR()
	}
	//否则平衡
	return at
}
