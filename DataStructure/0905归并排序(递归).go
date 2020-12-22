package stl

import "fmt"

//func (sl *SqList) MergeSort() {
//	sl.data=mergeSort(sl.data)
//}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	//拆分过程
	//递归调用，将arr一直往下裁成2块
	//递归左半部分0 --- mid-1 ,返回排好序的子序列
	l := mergeSort(arr[0:mid])
	//递归又半部分 mid -- len，返回排好序的子序列
	r := mergeSort(arr[mid:])
	//合并排好序的子序列
	return merge(l, r)
}

//接收2个子序列，返回一个排好序的完整序列
func merge(l, r []int) (result []int) {
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			result = append(result, l[i])
			i++
		} else {
			result = append(result, r[j])
			j++
		}
	}
	if i == len(l) {
		result = append(result, r[j:]...)
	}
	if j == len(r) {
		result = append(result, l[i:]...)
	}
	return
}

func main() {
	o := []int{2, 1, 3, 5, 6, 1, 8, 10, 7}
	fmt.Println(mergeSort(o))
}
