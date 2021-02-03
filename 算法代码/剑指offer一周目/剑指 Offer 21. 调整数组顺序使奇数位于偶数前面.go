package main

import "fmt"

//自己写的，遍历法，提交只有双10（不是双百(ㄒoㄒ)/~~）
func exchange(nums []int) []int {
	odd := make([]int, 0, len(nums))
	even := make([]int, 0, len(nums))
	for i, _ := range nums {
		if nums[i]%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	return append(odd, even...)
}

func exchange1(nums []int) []int {
	even := make([]int, 0, len(nums))
	for i := 0; i < len(nums); {
		if nums[i]%2 == 0 {
			even = append(even, nums[i])
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return append(nums, even...)
}

//标准答案双指针法，两边向中间
func exchange2(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 == 0 { //如果偶数在前
			if nums[j]%2 == 1 { //奇数在后，则互换位置
				t := nums[i]
				nums[i] = nums[j]
				nums[j] = t
				i++
				j--
			} else { //后一个也是偶数，则后一个向前找
				j--
			}
		} else { //如果奇数在前，
			if nums[j]%2 == 1 { //后一个也是奇数
				i++
			} else { //偶数在后
				i++
				j--
			}
		}
	}
	return nums
}

//结构优化1，时间反而多了
func exchange3(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 == 0 { //如果偶数在前
			if nums[j]%2 == 1 { //奇数在后，则互换位置
				t := nums[i]
				nums[i] = nums[j]
				nums[j] = t
				i++
			} //后一个也是偶数，则后一个向前找
			j--
		} else { //如果奇数在前，
			if nums[j]%2 == 0 { //后一个也是奇数
				j--
			} //偶数在后
			i++
		}
	}
	return nums
}

//使用位运算判断奇偶，时间花费更多了
func exchange4(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]&1 == 0 { //如果偶数在前
			if nums[j]&1 == 1 { //奇数在后，则互换位置
				t := nums[i]
				nums[i] = nums[j]
				nums[j] = t
				i++
			} //后一个也是偶数，则后一个向前找
			j--
		} else { //如果奇数在前，
			if nums[j]&1 == 0 { //偶数在后
				j--
			}
			i++ //后一个也是奇数
		}
	}
	return nums
}

//标准答案双指针法2，开头到结尾
func exchange5(nums []int) []int {
	f, s := 0, 0
	for s <= len(nums)-1 {
		if nums[s]%2 != 0 {
			t := nums[s]
			nums[s] = nums[f]
			nums[f] = t
			f++
		}
		s++
	}
	return nums
}
func main() {
	nums := []int{2, 16, 3, 5, 13, 1, 16, 1, 12, 18, 11, 8, 11, 11, 5, 1}
	fmt.Println(exchange5(nums))
}
