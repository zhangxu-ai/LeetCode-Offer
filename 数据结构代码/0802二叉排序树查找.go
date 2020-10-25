package main


//在有序二叉树中查找某个key并返回该节点
func (bn *BitNode)SearchByOrder(key interface{})(*BitNode,bool){
	if bn==nil{
		return bn,false
	}
	if bn.data.(int)==key.(int){
		return bn,true
	}else if key.(int)<bn.data.(int){
		return bn.lChild.SearchByOrder(key)
	}else {
		return bn.rChild.SearchByOrder(key)
	}
}

func (bn *BitNode) InsertByOrder(key interface{}) bool {
	if tn,ok:=bn.SearchByOrder(key);ok{
		return false
	}else {
		nn:=&BitNode{
			data:   key,
			lChild: nil,
			rChild: nil,
		}
		if tn==nil{
			tn=nn
		}else if key.(int)<tn.data.(int){
			tn.lChild=nn
		}else if key.(int)>tn.data.(int){
			tn.rChild=nn
		}
		return true
	}
}
