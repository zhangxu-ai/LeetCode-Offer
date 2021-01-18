package main

import (
	"container/list"
)

// 双端队列结构
type Deque struct {
	deque *list.List
}

// 初始化
func NewDeque() *Deque {
	l := list.New()
	return &Deque{l}
}

// 头部插入
func (dq *Deque) PushFront(el interface{}) {
	dq.deque.PushFront(el)
	return
}

// 尾部插入
func (dq *Deque) PushBack(el interface{}) {
	dq.deque.PushBack(el)
	return
}

// 查看头部元素的值
func (dq *Deque) Front() interface{} {
	if dq.deque.Front() == nil {
		return nil
	}
	return dq.deque.Front().Value
}

// 查看尾部元素的值
func (dq *Deque) Back() interface{} {
	if dq.deque.Back() == nil {
		return nil
	}
	return dq.deque.Back().Value
}

// 头部删除
func (dq *Deque) PopFront() {
	front := dq.deque.Front()
	dq.deque.Remove(front)
}

// 尾部删除
func (dq *Deque) PopBack() {
	back := dq.deque.Back()
	dq.deque.Remove(back)
}

// 队列长度
func (dq *Deque) Len() int {
	return dq.deque.Len()
}

// 队列是否为空
func (dq *Deque) IsEmpty() bool {
	return dq.deque.Len() == 0
}

// 定义单调队列
type MoQueue struct {
	data *Deque
}

// 初始化deque
func NewMoQueue() MoQueue {
	return MoQueue{
		data: NewDeque(),
	}
}

// 往单调队列里插入一个元素
func (Mo *MoQueue) Push(x int) {
	for Mo.data.IsEmpty() == false && x > Mo.data.Back().(int) {
		Mo.data.PopBack()
	}
	Mo.data.PushBack(x)
}

// 单调队列的最大值
func (Mo *MoQueue) Max() int {
	return Mo.data.Front().(int)
}

// 单调队列pop一个元素
func (Mo *MoQueue) Pop(x int) {
	if Mo.data.IsEmpty() == false && Mo.data.Front() == x {
		Mo.data.PopFront()
	}
}

// 最大滑动窗口
func maxSlidingWindow3(nums []int, k int) []int {
	result := make([]int, 0)

	if len(nums) == 0 {
		return result
	}

	window := NewMoQueue()

	for i := 0; i < len(nums); i++ {

		if i < k-1 { // 将k-1部分先入队
			window.Push(nums[i])
		} else {
			window.Push(nums[i])
			result = append(result, window.Max())
			window.Pop(nums[i-k+1]) // 将窗口最左边元素弹出
		}

	}
	return result

}
