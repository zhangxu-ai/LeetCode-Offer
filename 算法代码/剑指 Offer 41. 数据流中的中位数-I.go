package main

import (
	"fmt"
	"sort"
)

//自己的想法，插入时二分插入，使之一直是排好序的队列
//6,7
//借用tsj的话：这是什么电子垃圾啊
type MedianFinder struct {
	data []int
}

/** initialize your data structure here. */
func ConstructorMedian() MedianFinder {
	return MedianFinder{
		data: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	l := len(this.data)
	if l == 0 {
		this.data = append(this.data, num)
		return
	} else if l == 1 {
		if this.data[0] <= num {
			this.data = append(this.data, num)
		} else {
			this.data = append([]int{num}, this.data...)
		}
		return
	}
	//二分
	low, high := 0, l-1
	for high >= low {
		mid := (low + high) / 2
		if mid == high && this.data[mid] <= num {
			this.data = append(this.data[:mid+1], append([]int{num}, this.data[mid+1:]...)...)
			//tmp:=make([]int,l-mid-1)
			//copy(tmp,this.data[mid+1:])
			//this.data = append(this.data[:mid+1],num)
			//this.data = append(this.data, tmp...)
			return
		} else if low == 0 && mid == low && this.data[mid] >= num {
			//this.data = append(this.data[:mid+1], append([]int{num}, this.data[mid+1:]...)... )
			this.data = append([]int{num}, this.data...)
			return
		} else if this.data[mid] <= num && this.data[mid+1] >= num {
			this.data = append(this.data[:mid+1], append([]int{num}, this.data[mid+1:]...)...)
			//tmp:=make([]int,l-mid-1)
			//copy(tmp,this.data[mid+1:])
			//this.data = append(this.data[:mid+1],num)
			//this.data = append(this.data, tmp...)
			return
		} else if this.data[mid] > num {
			high = mid - 1
		} else if this.data[mid] < num {
			low = mid + 1
		}
	}
}

//调用sort库函数排序，好家伙直接超时
func (this *MedianFinder) AddNum1(num int) {
	this.data = append(this.data, num)
	sort.Ints(this.data)
}

func (this *MedianFinder) FindMedian() float64 {
	l := len(this.data)
	if l == 0 {
		return 0.0
	}
	if l%2 == 0 {
		i := this.data[(l-1)/2] + this.data[l/2]
		return float64(i) / 2
	} else {
		return float64(this.data[l/2])
	}
}

func main() {
	a := ConstructorMedian()
	a.AddNum1(6)
	fmt.Println(a.FindMedian())
	a.AddNum1(10)
	fmt.Println(a.FindMedian())
	a.AddNum1(2)
	fmt.Println(a.FindMedian())
	a.AddNum1(6)
	fmt.Println(a.FindMedian())
	a.AddNum1(5)
	fmt.Println(a.FindMedian())
	a.AddNum1(0)
	fmt.Println(a.FindMedian())
	a.AddNum1(6)
	fmt.Println(a.FindMedian())
	a.AddNum1(3)
	fmt.Println(a.FindMedian())
	a.AddNum1(1)
	fmt.Println(a.FindMedian())
	a.AddNum1(0)
	fmt.Println(a.FindMedian())
	a.AddNum1(0)
	fmt.Println(a.FindMedian())

}
