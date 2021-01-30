# 剑指 Offer 30. 包含min函数的栈

## 题目

定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:

```go
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   //--> 返回 -3.
minStack.pop();
minStack.top();   //   --> 返回 0.
minStack.min();   //--> 返回 -2.
```

## 思考

和上一题有个共同点，可以使用辅助栈，A存真正数据，B存当前最小值。
当push时，先push到a中，然后将x和B的top对比。如果x大，则不放，因为即使x被push出，最小的依然是b的top，如果小于等于，则push
在pop时，如果b的top和a的top一样大，则a和b一起pop出；否则只pop a 的

## 代码

```go
type MinStack struct {
	data []int
	min []int
}

func ConstructorMS() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{},
	}
}


func (s *MinStack) Push(x int)  {
	s.data = append(s.data, x)
	if len(s.min)==0||s.min[len(s.min)-1]>=x{
		s.min = append(s.min, x)
	}
}


func (s *MinStack) Pop()  {
	l:= len(s.data)
	if l==0{
		return
	}
	if s.data[l-1]==s.min[len(s.min)-1]{
		s.min=s.min[:len(s.min)-1]
	}
	s.data=s.data[:l-1]
}


func (s *MinStack) Top() int {
	l:= len(s.data)
	if l==0{
		return -1
	}
	return s.data[l-1]
}

func (s *MinStack) Min() int {
	l:= len(s.min)
	if l==0{
		return -1
	}
	return s.min[l-1]
}

```
