package main

import "fmt"

//优化后的暴力法
var ns []int
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums)==0||k==1{
		return nums
	}
	ns=nums
	res:=make([]int,0,len(nums)-k+1)
	i,j,maxIndex:=0,k-1,0
	for j<=len(nums)-1{
		if i==0{
			maxIndex=max(i,k)
		}else{
			if nums[j]>=nums[maxIndex]{
				maxIndex=j
			}else if maxIndex==i-1{
				maxIndex=max(i,k)
			}
		}
		res = append(res, nums[maxIndex])
		i++
		j++
	}
	return res
}

func max(start, len int) int {
	maxIndex:=start
	for i := start; i < start+len; i++ {
		if ns[i]>=ns[maxIndex]{
			maxIndex=i
		}
	}
	return maxIndex
}

func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums)==0||k==1{
		return nums
	}
	res:=make([]int,0,len(nums))
	//初始化单调队列
	deque:=make([]int,0,k)
	//将前k个入队，即第一个窗口
	for i := 0; i < k; i++ {
		for len(deque)!=0 &&deque[len(deque)-1]<nums[i] {
			deque=deque[:len(deque)-1]
		}
		deque = append(deque, nums[i])
	}
	res = append(res, deque[0])
	
	//开始滑动窗口
	for i := 1; i <= len(nums)-k ; i++ {
		//把新的加进去
		for len(deque)!=0 &&deque[len(deque)-1]<nums[i+k-1] {
			deque=deque[:len(deque)-1]
		}
		deque = append(deque, nums[i+k-1])
		//把前一个去掉
		if deque[0]==nums[i-1]{
			deque=deque[1:]
		}
		res = append(res, deque[0])
	}
	return res
}

func main() {
	a:=[]int{1,3,-1,-3,5,3,6,7}
	k:=3
	fmt.Println(maxSlidingWindow1(a,k))
}
