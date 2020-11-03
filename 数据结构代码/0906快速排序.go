package main

import "fmt"

func Swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
func QuickSort(arr []int, l, h int) {
	if l < h {
		pivot := Partition2(arr, l, h)
		//对左半区进行快排
		QuickSort(arr, l, pivot-1)
		//对右半区进行快排
		QuickSort(arr, pivot+1, h)
	}
}

//i和j一个向右，一个向左
func Partition(arr []int, l, h int) int {
	//用子表的第一个记录作为key
	//因为后面先扫描大的，所以会把这个key放到后面，
	//在扫描小的时候，也可能把key放到前面去
	//并且每一轮循环中，l和h必定有一个指向的是这个key
	//实际上就是这个key和其他元素换来换去
	//直到l和h重合，即扫描完所有的元素
	pivotKey := arr[l]
	for l < h {
		//先向从大向小扫描，直到找到比pivotKey小的值
		for l < h && arr[h] >= pivotKey {
			h--
		}
		//将小的值放到前面，此时最小值可能是
		Swap(arr, l, h)
		//第一步扫面完后再从小向大扫描，直到找到比pivotKey大的值
		for l < h && arr[l] <= pivotKey {
			l++
		}
		//将大的值放到后面
		Swap(arr, l, h)
		//继续下一轮循环
	}
	//上面的循环会在l==h时退出，此时比pivot小的值都在l的左边，大的都在右边
	return l
}

//i和j变换的方向相同，都是向右
func Partition2(arr []int, l, h int) int {
	//取最高位为基准
	k := arr[h]
	i := l - 1
	//j初始时从l出发，向右递增；当遇到比k基准值小的时候，就将i向右移动一位，
	//然后交换j和l的数据，这样就把小于k的数据放到了数组的左边，右边都是大的数据
	for j := l; j < h; j++ {
		if arr[j] < k {
			i++
			Swap(arr, i, j)
		}
	}
	//直到j==h,这个时候把基准值和i右边的一位交换位置
	//因为i左边都是小的，右边是大的，而基准值本身就在最右边，所以把一个大的换过去
	Swap(arr, i+1, h)
	//将k现在的位置返回，当做下一次的基准点位置
	return i + 1
}

func main() {
	o := []int{2, 1, 3, 5, 6, 1, 8, 10, 7}
	//fmt.Println(Partition(o,0,len(o)-1))
	QuickSort(o, 0, len(o)-1)
	fmt.Println(o)
}
