package main

//has表法，空间不好
//72,22
func singleNumbers(nums []int) []int {
	l := len(nums)

	tmp := make(map[int]int8, l/2+1)

	for i := 0; i < l; i++ {
		if _, ok := tmp[nums[i]]; ok {
			delete(tmp, nums[i])
		} else {
			tmp[nums[i]] = 1
		}
	}
	res := make([]int, 0, 2)
	for k := range tmp {
		res = append(res, k)
	}
	return res
}

//亦或法
//开始先到了遍历一遍，得到a亦或b，但是后续不知道如何处理，只能看题解了
//72,99
func singleNumbers1(nums []int) []int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}
	div := 1
	//找到ret最低非0位
	//方法1
	for (div & ret) == 0 {
		div <<= 1
	}
	//方法2
	//负数为原来的反码然后+1,
	//div = ret & (-ret)
	//再次遍历nums
	//根据div进行分组
	/*①相同的数字一定在一组
	②a和b在不同的分组。组内异或即可找出a和b
	(其实只要找出a或b)
	*/
	//a,b:=0,0
	//for i := 0; i < len(nums); i++ {
	//	if (div&nums[i])==0{
	//		a^=nums[i]
	//	}else{
	//		b^=nums[i]
	//	}
	//}
	//return []int{a,b}

	//只找一个就行
	a := 0
	for i := 0; i < len(nums); i++ {
		if (div & nums[i]) == 0 {
			a ^= nums[i]
		}
	}
	ret ^= a
	return []int{a, ret}

}
