package main

//69,60
type MinStack2 struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor3() MinStack2 {
	return MinStack2{
		data: []int{},
		min:  []int{},
	}
}

func (this *MinStack2) Push(x int) {
	this.data = append(this.data, x)
	l := len(this.min)
	if l > 0 {
		if this.min[l-1] >= x {
			this.min = append(this.min, x)
		}
	} else {
		this.min = append(this.min, x)
	}
}

func (this *MinStack2) Pop() {
	l := len(this.data)
	if l == 0 {
		return
	}
	if this.data[l-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.data = this.data[:l-1]
}

func (this *MinStack2) Top() int {
	l := len(this.data)
	if l == 0 {
		return 0
	}
	return this.data[l-1]
}

func (this *MinStack2) Min() int {
	l := len(this.min)
	if l == 0 {
		return 0
	}
	return this.min[l-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
