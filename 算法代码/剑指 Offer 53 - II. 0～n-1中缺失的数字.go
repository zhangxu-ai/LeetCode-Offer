package main

//二分
//66,58
func missingNumber(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	left, right, mid := 0, l-1, 0
	for left <= right {
		mid = left + ((right - left) >> 1)
		//不可能出现nums[mid]<mid的情况
		if nums[mid] == mid {
			left = mid + 1
		} else if nums[mid] > mid {
			right = mid - 1
		}
	}
	return left
}

//遍历
//11,100
func missingNumber1(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	for i := 0; i < l; i++ {
		if i == nums[i] {
			continue
		}
		return i
	}

	return nums[l-1] + 1
}

//遍历+亦或
//11,58
/*
异或运算的性质：
1．两个数相同，则异或运算结果为 0；
2．任何数和 0 异或的结果还是自身。
3. a^b^b=a
本题中下标范围：[0,n-1]、数值范围[0,n]，
把下标和数值依次进行异或，最终的结构即为只出现一次的值，即缺失的值。
*/
func missingNumber2(nums []int) int {
	l, r := len(nums), 0
	if l == 0 {
		return 0
	}
	for i := 0; i < l; i++ {
		r ^= i
		r ^= nums[i]
	}
	//最后要和l再亦或一次，因为上面遍历到最后时多了一个nums[l-1]
	//因为nums里面缺了一个，所以nums[l-1]=l
	//如果缺的是最后一个，则r此时=0，和l亦或变成了l，妙
	return r ^ l
}
