package main

import "fmt"

//从尾部遍历nums，将此元素和dp里的元素比较，
//因为dp是排好序的，所以只要找到第一个比他小的，
//后面都是比他小，就全都是它的逆序
//然后再将此元素添加到dp中，并排序
//超时，电子垃圾
var dp []int

func reversePairs(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}
	var res = 0
	dp = make([]int, 0, l)
	for i := l - 1; i >= 0; i-- {
		if i == l-1 {
			dp = append(dp, nums[i])
		} else {
			for j := len(dp) - 1; j >= 0; j-- {
				if dp[j] >= nums[i] {
					continue
				} else {
					res += j + 1
					break
				}
			}
			AddNum(nums[i])
		}
	}
	return res
}

//使用二分往dp里面插入数字到正确位置
func AddNum(num int) {
	l := len(dp)
	if l == 0 {
		dp = append(dp, num)
		return
	} else if l == 1 {
		if dp[0] <= num {
			dp = append(dp, num)
		} else {
			dp = append([]int{num}, dp...)
		}
		return
	}
	//二分
	low, high := 0, l-1
	for high >= low {
		mid := (low + high) / 2
		if mid == high && dp[mid] <= num {
			dp = append(dp[:mid+1], append([]int{num}, dp[mid+1:]...)...)
			//tmp:=make([]int,l-mid-1)
			//copy(tmp,this.data[mid+1:])
			//this.data = append(this.data[:mid+1],num)
			//this.data = append(this.data, tmp...)
			return
		} else if low == 0 && mid == low && dp[mid] >= num {
			//this.data = append(this.data[:mid+1], append([]int{num}, this.data[mid+1:]...)... )
			dp = append([]int{num}, dp...)
			return
		} else if dp[mid] <= num && dp[mid+1] >= num {
			dp = append(dp[:mid+1], append([]int{num}, dp[mid+1:]...)...)
			//tmp:=make([]int,l-mid-1)
			//copy(tmp,this.data[mid+1:])
			//this.data = append(this.data[:mid+1],num)
			//this.data = append(this.data, tmp...)
			return
		} else if dp[mid] > num {
			high = mid - 1
		} else if dp[mid] < num {
			low = mid + 1
		}
	}
	return
}

//归并排序法
//太难先不看
func main() {
	a := []int{7, 5, 6, 4}
	fmt.Println(reversePairs(a))
}
