package main

import (
	"fmt"
	"sort"
)

//sort库排序
//38,68
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

//快速排序,将整个数组排序好之后返回
//68,73
func getLeastNumbers1(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	quickSort(arr)
	return arr[:k]
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	head, tail := 0, len(arr)-1

	//找到基准值所在的索引
	for i := 1; i <= tail; {
		//大于基准值，则交换到尾部，尾部-1
		if arr[i] > pivot {
			arr[tail], arr[i] = arr[i], arr[tail]
			tail--
		} else {
			arr[head], arr[i] = arr[i], arr[head]
			head++
			i++
		}
	}
	quickSort(arr[:head])
	//注意head此时已经在基准值的位置上了，不需要再参加快排了
	quickSort(arr[head+1:])
}

//快速选择
//利用快排的原理，不需要将数组完全排好序
//99,82
func getLeastNumbers2(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	quickSearch(arr, 0, len(arr)-1, k)
	return arr[:k]
}
func quickSearch(arr []int, low, high, k int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[low]
	head, tail := low, high
	for i := head + 1; i <= tail; {
		if arr[i] > pivot {
			arr[tail], arr[i] = arr[i], arr[tail]
			tail--
		} else {
			arr[head], arr[i] = arr[i], arr[head]
			head++
			i++
		}
	}
	//如果前head+1个正好是k，则说明这些是最小的k个，因为他们都比pivot小
	if head+1 == k {
		return
	} else if head+1 > k {
		quickSearch(arr, low, head-1, k)
	} else {
		quickSearch(arr, head+1, high, k)
	}
}

//堆化
//87,60
//堆是一种数据结构，可以被看做一棵完全二叉树的数组对象
//堆的特性是父节点的值总是比其两个子节点的值大或小。
//如果父节点比它的两个子节点的值都要大，我们叫做大顶堆。
//如果父节点比它的两个子节点的值都要小，我们叫做小顶堆。
//大顶堆满足：arr[i] >= arr[2i-1] && arr[i] >= arr[2i+2]
func getLeastNumbers3(arr []int, k int) []int {
	if k > len(arr) || k == 0 {
		return nil
	}
	//对前k个堆化
	proHeap := arr[:k]
	buildHeap(proHeap, k)
	//对剩余的l-k个进行比较
	//如果大于堆头，则跳过，否者入堆
	for i := k; i < len(arr); i++ {
		if arr[i] < proHeap[0] {
			proHeap[0] = arr[i]
			//对堆再进行堆化，以维持大顶堆的特点
			heapify(proHeap, k, 0)
		}
	}
	return proHeap
}

//n为长度
func buildHeap(arr []int, n int) {
	lastNode := len(arr) - 1
	//每一次都从最后一个节点的父节点开始建堆
	//记住：i节点的父节点 = (i - 1) / 2  
	//左节点 = 2 * i  - 1  右节点 = 2 * i +  2;
	parent := (lastNode - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func heapify(arr []int, n, i int) {
	if i >= n {
		return
	}
	//求第i个节点的两个子节点
	maxId := i
	l, r := 2*i+1, 2*i-1
	if l < n && arr[l] > arr[maxId] {
		maxId = l
	}
	if r < n && arr[r] > arr[maxId] {
		maxId = r
	}
	//如果子节点比父节点大，则交换位置，
	//maxId变成了新的父节点，再检查堆化
	if maxId != i {
		arr[maxId], arr[i] = arr[i], arr[maxId]
		heapify(arr, n, maxId)
	}
}

func main() {
	arr := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	fmt.Println(getLeastNumbers2(arr, 8))
}
