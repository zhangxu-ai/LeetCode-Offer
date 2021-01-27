# 剑指 Offer 59 - II. 队列的最大值

## 题目

请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：
输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]

示例 2：
输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]

限制：

1. 1 <= push_back,pop_front,max_value的总操作数 <= 10000
2. 1 <= value <= 10^5

## 思考

和上一题的min栈相似，只不过这次是求max，并且是队列，即先进先出
当append时，先append到a中，然后将x和B的head对比。如果x大于等于，则把b的前面的头部都丢掉，然后把x变成新的头部。因为在deleteHead时，直到x时，最大的都是x；如果x小，则往后遍历，直到有比x小的值，把这个值之后的都丢掉，然后把x放到这个值的位置上。

## 代码

```go
type MaxQueue struct {
	data []int
	max []int
}

func ConstructorMQ() MaxQueue {
	return MaxQueue{
		data: []int{},
		max:  []int{},
	}
}


func (q *MaxQueue) Max_value() int {
	if len(q.max)==0{
		return -1
	}else{
		return q.max[0]
	}
}

func (q *MaxQueue) Push_back(value int)  {
	q.data = append(q.data, value)
	if len(q.max)==0{
		q.max = append(q.max, value)
	}else {
		i:=0
		for i = 0; i < len(q.max)&&q.max[i]>=value; i++ {
		}
		q.max=q.max[:i]
        q.max=append(q.max,value)
	}
}


func (q *MaxQueue) Pop_front() int {
	if len(q.data)==0{
		return -1
	}
	t:=q.data[0]
	q.data=q.data[1:]
	if t==q.max[0]{
		q.max=q.max[1:]
	}
	return t
}


```
