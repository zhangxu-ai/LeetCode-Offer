package main

import "fmt"

// in 入队数据  out 出队数据
type CQueue2 struct {
	in  []int
	out []int
}

func Constructor2() CQueue2 {
	return CQueue2{}
}

// 队尾入队
func (this *CQueue2) AppendTail(value int) {
	this.in = append(this.in, value)
}

// 出队
func (this *CQueue2) DeleteHead() int {
	// 空
	if len(this.out) == 0 {
		for i := 0; i < len(this.in); i++ {
			// 依次放在出队数据
			this.out = append(this.out, this.in[i])
		}
		// 清空入队队列
		this.in = []int{}
	}

	if len(this.out) == 0 {
		return -1
	}
	outEleme := this.out[0]
	// 更新出队队列
	this.out = this.out[1:]
	return outEleme
}

func main() {
	c := Constructor2()
	c.AppendTail(3)
	fmt.Println(c.DeleteHead(), c.DeleteHead(), c.DeleteHead())
}
