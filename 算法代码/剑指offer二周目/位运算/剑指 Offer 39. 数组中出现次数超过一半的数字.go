package main

//只写位运算的方法
func majorityElement(nums []int) int {
	//注意res一定要是int32
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
			//fmt.Printf("num:%b t: %b\n",nums[i2],t)
			if int32(nums[i2])&t!=0{
				c1++
			}
			//如果第i位里，1占多数，则res上第i位也为1
			if c1>l/2{
				//1向左移动i位，结果就是第i位为1，其余位还是0
				//然后加到res上。因为res初始为0，所以res的第i位也是1了
				res |= t
				//fmt.Printf("i:%d res:%t %v %b\n",i,res,res,res)
				break
			}
		}
	}

	return int(res)
}
