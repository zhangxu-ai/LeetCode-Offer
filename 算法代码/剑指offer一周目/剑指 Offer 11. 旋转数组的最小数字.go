package main

import "fmt"

//都是使用双指针，但是第一次想要根据头、中、尾的大小关系，
//确定要走的方向，进而从中间向一边遍历比较
//但是在在面对如[3,3,3,3,1,3]和[3,1,3,3,3,3,3]这种情况时，
//就没办法了，所以只能接收从一边遍历的事实
//func minArray(ns []int) int {
//	l:=len(ns)
//	if l==1{return ns[0]}
//	if l==2{
//		if ns[0]<ns[1]{
//			return ns[0]
//		}
//		return ns[1]
//	}
//	f,s:=l/2-1,l/2
//	//顺序排列
//	if ns[0]<ns[l-1]{
//		return ns[0]
//	}else if ns[s]<=ns[l-1]{ //旋转了,且s在右半部分
//		for ns[f]<=ns[s]&&f>0{
//			f--
//			s--
//		}
//		return ns[s]
//	}else {
//		for ns[f]<=ns[s]&&s<l-1{
//			f++
//			s++
//		}
//		return ns[s]
//	}
//}
//
func minArray(ns []int) int {
	l := len(ns)
	if l < 2 || ns[0] < ns[l-1] {
		return ns[0]
	}
	for ns[l-2] <= ns[l-1] && l > 2 {
		l--
	}
	return ns[l-1]
}
func main() {
	a := []int{2, 2, 2, 0, 1}
	fmt.Println(minArray(a))
}
