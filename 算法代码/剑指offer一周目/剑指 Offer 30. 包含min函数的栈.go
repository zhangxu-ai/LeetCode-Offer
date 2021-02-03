package main

import (
	"fmt"
)

//100,68
type MinStack struct {
	data    []int
	min     int
	isMined bool
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: []int{}, min: 0, isMined: false}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.data) == 0 {
		this.min = x
		this.isMined = true
	} else if this.isMined {
		if this.min > x {
			this.min = x
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	t := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if this.isMined && t == this.min {
		this.isMined = false
		_ = this.Min()
	}
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	if len(this.data) == 0 {
		return 0
	}
	if this.isMined {
		return this.min
	}
	min := this.data[0]
	for i := 0; i < len(this.data); i++ {
		if min > this.data[i] {
			min = this.data[i]
		}
	}
	this.min = min
	this.isMined = true
	return min
}
func main() {
	a := Constructor()
	a.Push(2)
	a.Push(0)
	a.Push(3)
	a.Push(0)
	fmt.Println(a.data)
	fmt.Println(a.Min())
	fmt.Println(a.data)
	a.Pop()
	fmt.Println(a.data)

	fmt.Println(a.Min())
	a.Pop()
	fmt.Println(a.data)

	fmt.Println(a.Min())
	a.Pop()
	fmt.Println(a.data)
	fmt.Println(a.Min())
	fmt.Println(a.data)
}
