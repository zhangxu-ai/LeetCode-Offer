package stl

//在有序二叉树中查找某个key并返回该节点
func (bn *BitNode) SearchByOrder(key interface{}) (*BitNode, bool) {
	if bn == nil {
		return bn, false
	}
	if bn.data.(int) == key.(int) {
		return bn, true
	} else if key.(int) < bn.data.(int) {
		return bn.lChild.SearchByOrder(key)
	} else {
		return bn.rChild.SearchByOrder(key)
	}
}

func (bn *BitNode) InsertByOrder(key interface{}) bool {
	if tn, ok := bn.SearchByOrder(key); ok {
		return false
	} else {
		nn := &BitNode{
			data:   key,
			lChild: nil,
			rChild: nil,
		}
		if tn == nil {
			tn = nn
		} else if key.(int) < tn.data.(int) {
			tn.lChild = nn
		} else if key.(int) > tn.data.(int) {
			tn.rChild = nn
		}
		return true
	}
}

func (bn *BitNode) DeleteBST(key interface{}) bool {
	if tn, ok := bn.SearchByOrder(key); !ok {
		return false
	} else {
		//右子树为空，用左子树替换
		if tn.rChild == nil {
			tn = tn.lChild
		} else if tn.lChild == nil {
			tn = tn.rChild //左子树为空，用右子树替换
		} else {
			q := tn
			s := tn.lChild
			//左右子树都不为空，则找到前继节点,就是其左子树的最右节点
			for s.rChild != nil {
				q = s
				s = s.rChild
			}
			tn.data = s.data
			//如果其左子树有右子树，
			//则接上这个前继节点的双亲节点到此前继节点的左子树
			if q != tn {
				q.rChild = s.lChild
			} else {
				q.lChild = s.lChild
			}
		}
	}
	return true
}
