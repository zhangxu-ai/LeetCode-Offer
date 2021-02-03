package main

//哈希表法
//66,46
func singleNumber(nums []int) int {
	l := len(nums)

	tmp := make(map[int]int8, l/3+1)

	for i := 0; i < l; i++ {
		if v, ok := tmp[nums[i]]; ok {
			if v == 2 {
				delete(tmp, nums[i])
			} else {
				tmp[nums[i]] = 2
			}
		} else {
			tmp[nums[i]] = 1
		}
	}
	for k := range tmp {
		return k
	}
	return 0
}

//使用位运算的性质来算
//根据题意我们可以找出这么一个规律：
//假设不存在这个 single number，那么 nums 中所有元素的二进制形式加起来之后，每一位必然都可以被 3 整除。
//因为每个数字都出现了三次，那么它们的二进制形式每一位也都出现了三次，那加起来之后每一位当然可以被 3 整除了。
//现在加上这个 single number，我们可以进一步想到：
//同样把 nums 中所有元素的二进制形式加起来，这时候就不是每一位都能被 3 整除了，
//因为混入了一个只出现了一次的元素。
//所以我们只需要看加起来之后的和的哪一位不能被 3 整除，
//就说明这个 single number 的二进制形式在这一位是 1。
//然后把这些 single number 是 1 的位转换成十进制再加起来，就是 single number 的十进制形式了。
//
func singleNumber1(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		cnt := 0
		bit := 1 << i
		for i := 0; i < len(nums); i++ {
			if nums[i]&bit != 0 {
				cnt++
			}
		}
		//如果此位的3余数不为0，说明是res这一位是1
		if cnt%3 != 0 {
			//使用|=将此位置为1
			res |= bit
		}
	}
	return res
}
