# 剑指 Offer 09. 用两个栈实现队列

## 题目

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

示例 1：

输入：
`["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]`  
输出：`[null,null,3,-1]`  

示例 2：

输入：
`["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]  
[[],[],[5],[2],[],[]]`
输出：`[null,-1,null,null,5,2]`

提示：  
1 <= values <= 10000  
最多会对appendTail、deleteHead 进行10000次调用

## 思考

两个栈，一个负责 appendTail ，一个负责 deleteHead。
当append时，正常append到A栈；delete是，先将a的所有数据pop出去，然后push到B里面，
这个时候A的最后也就是最先放入的头部变成了B的尾部，就可以最先删除了

使用数组来模拟栈

## 代码

```go
type CQueue struct {
	a []int
	b []int
}

func Constructor() CQueue {
	return CQueue{
		a: []int{},
		b: []int{},
	}
}

func (q *CQueue) AppendTail(value int)  {
	q.a = append(q.a, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.a)==0&&len(q.b)==0{
		return -1
	}else if len(q.a)!=0&&len(q.b)==0{
		//模拟a出栈，b入栈过程
		for len(q.a)>0 {
			q.b = append(q.b, q.a[len(q.a)-1])
			q.a=q.a[:len(q.a)-1]
		}
	}
	t:=q.b[len(q.b)-1]
	q.b=q.b[:len(q.b)-1]
	return t
}```
