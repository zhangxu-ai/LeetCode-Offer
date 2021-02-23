package main

import "fmt"

func serialize(root *TreeNode) []int {
	if root==nil{
		return nil
	}
	var (
		res  []int
		tmp []*TreeNode
	)
	tmp = append(tmp, root)
	for i := 0; i < len(tmp); i++ {
		if tmp[i]==nil{
			res = append(res, -999)
		}else{
			res = append(res, tmp[i].Val)
			tmp = append(tmp, tmp[i].Left)
			tmp = append(tmp, tmp[i].Right)
		}
	}
	return res
}

func deserialize(a []int)*TreeNode  {
	if len(a)==0{
		return nil
	}
	var tmp []*TreeNode
	root:=&TreeNode{
		Val: a[0],
	}
	tmp = append(tmp, root)
	for i,j := 1,0; i < len(a); j++{
		t:=tmp[j]
		if a[i]!=-999{
			n:=new(TreeNode)
			n.Val=a[i]
			t.Left=n
			tmp = append(tmp, n)
		}
		i++
		if a[i]!=-999{
			n:=new(TreeNode)
			n.Val=a[i]
			t.Right=n
			tmp = append(tmp, n)
		}
		i++
	}
	return tmp[0]
}

func main() {
	a:=[]int{3,9,20,-999,-999,15,7}
	fmt.Println(serialize(deserialize(a)))

}
