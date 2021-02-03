package main

import (
	"fmt"
	"sort"
)

//38，6.1
func majorityElement(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	tmp := make(map[int]int, len(nums))
	for i := 0; i < l; i++ {
		if c, ok := tmp[nums[i]]; ok {
			if c >= l/2 {
				return nums[i]
			} else {
				tmp[nums[i]] = c + 1
			}
		} else {
			tmp[nums[i]] = 1
		}
	}
	return nums[0]
}

//数组排序法： 将数组 nums 排序，数组中点的元素 一定为众数。
//38,96
func majorityElement1(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	sort.Ints(nums)
	return nums[l/2]
}

//摩尔投票法1
//83,55
func majorityElement2(nums []int) int {
	x, v := 0, 0
	for i := 0; i < len(nums); i++ {
		if v == 0 {
			x = nums[i]
		}
		if nums[i] == x {
			v += 1
		} else {
			v -= 1
		}
	}
	return x
}

//位运算法
/*
给出输入[1,2,3,3,3]，将这几个数字写成二进制数(整型32位，这里只写出后面4位)[0001, 0010, 0011, 0011, 0011]。
在这4位二进制数字中。各自的最高位分别为[0，0，0，0]，则最高位出现最多的数字是二进制0。
同理，高二位出现最多的二进制数字为0，高三位出现最多的二进制数字为1，最低位出现最多的二进制数字为1。
则4位出现最多的二进制数字拼起来为0011，值为3，即结果。
*/
//因为负数是补码形式，所以造成结果都很大，比如-1为4294967295，暂时找不到解决方法
func majorityElement3(nums []int) int {
	res,l:=int32(0),len(nums)
	if l==0{
		return 0
	}
	if l<=2{
		return nums[0]
	}
	for i := 0; i < 32; i++ {
		c1:=0
		t:=int32(1<<i)
		for i2 := 0; i2 < l; i2++ {
			//fmt.Printf("%b\n%b\n",uint32(nums[i2]),t)
			//判断第i位是1还是0
			fmt.Printf("num:%b t: %b\n",nums[i2],t)
			if int32(nums[i2])&t!=0{
				c1++
			}
			//如果第i位里，1占多数，则res上第i位也为1
			if c1>l/2{
				//1向左移动i位，结果就是第i位为1，其余位还是0
				//然后加到res上。因为res初始为0，所以res的第i位也是1了
				res |= t
				fmt.Printf("i:%d res:%t %v %b\n",i,res,res,res)
				break
			}
		}
	}

	return int(res)
}
//
func main() {
	a := []int{-1, -1, 2147483647}
	fmt.Println(majorityElement3(a))
}
