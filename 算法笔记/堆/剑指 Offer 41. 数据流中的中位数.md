# 剑指 Offer 41. 数据流中的中位数

## 题目

如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

`void addNum(int num)` - 从数据流中添加一个整数到数据结构中。
`double findMedian()` - 返回目前所有元素的中位数。
示例 1：

输入：
`["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]`
`[[],[1],[2],[],[3],[]]`
输出：`[null,null,null,1.50000,null,2.00000]`
示例 2：

输入：
`["MedianFinder","addNum","findMedian","addNum","findMedian"]`
`[[],[2],[],[3],[]]`
输出：`[null,null,2.00000,null,2.50000]`

## 思考

普通方法就是每次加入的时候都排序，然后中位数直接返回中间一个或者2个的平均值。
使用大小顶堆的方法，小的一半使用大顶堆，大的一半使用小顶堆，则2个堆的顶部一定是数据流的中间两个数。其中，如果是奇数个，则自接使用大顶堆的第一个，如果是偶数个，则返回2个堆的顶部平均值。
堆是动态调整的，非常适合在数据流中使用

当插入一个数据时，如果当前maxHeap(存较小部分的数据)的长度等于minHeap的长度，则先把num存到minHeap最后，然后重新堆化。之后min的头部存的是最小部分，然后再把头部抛出，给max，max重新堆化

## 代码

```go
func heapMaxDown(arr []int,id int) {
	left:=id*2+1
	right:=left+1
	max:=id
	if left<len(arr)&&arr[left]>arr[max]{
		max=left
	}
	if right<len(arr)&&arr[right]>arr[max]{
		max=right
	}
	if max!=id{
		arr[max],arr[id]=arr[id],arr[max]
		heapMaxDown(arr,max)
	}
}

func heapMinDown(arr []int,id int) {
	left:=id*2+1
	right:=left+1
	min :=id
	if left<len(arr)&&arr[left]<arr[min]{
		min =left
	}
	if right<len(arr)&&arr[right]<arr[min]{
		min =right
	}
	if min !=id{
		arr[min],arr[id]=arr[id],arr[min]
		heapMinDown(arr, min)
	}
}
func buildHeap(arr []int,isMax bool)  {
	l:= len(arr)
	lastNode:=l-1
	lastParent:=(lastNode-1)/2
	if isMax{
		for i := lastParent; i >=0 ; i-- {
			heapMaxDown(arr,i)
		}
	}else{
		for i := lastParent; i >=0 ; i-- {
			heapMinDown(arr,i)
		}
	}
}
type MedianFinder struct {
	maxHeap []int
	minHeap []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: []int{},
		minHeap: []int{},
	}
}

func (f *MedianFinder) AddNum(num int)  {
	if len(f.maxHeap)==0{
		f.maxHeap = append(f.maxHeap, num)
		return
	}

	if len(f.maxHeap)== len(f.minHeap){
		f.minHeap = append(f.minHeap, num)
		buildHeap(f.minHeap,false)
		f.maxHeap = append(f.maxHeap, f.minHeap[0])
		buildHeap(f.maxHeap,true)
		f.minHeap=f.minHeap[1:]
		buildHeap(f.minHeap,false)
	}else{
		f.maxHeap = append(f.maxHeap, num)
		buildHeap(f.maxHeap,true)
		f.minHeap = append(f.minHeap, f.maxHeap[0])
		buildHeap(f.minHeap,false)
		f.maxHeap=f.maxHeap[1:]
		buildHeap(f.maxHeap,true)

	}
}

func (f *MedianFinder) FindMedian() float64 {
	if len(f.maxHeap)==0{
		return 0.0
	}
	if len(f.maxHeap)==len(f.minHeap){
		return (float64(f.maxHeap[0])+float64(f.minHeap[0]))/2
	}else{
		return float64(f.maxHeap[0])
	}
}
```
