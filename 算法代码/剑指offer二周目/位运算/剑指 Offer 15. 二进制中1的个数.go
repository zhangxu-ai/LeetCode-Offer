package main

import "fmt"

func hammingWeight(num uint32) int {
	res:=0
	for num != 0 {
		if num&1==1{
			res++
		}
		num>>=1
	}
	return res
}

//优化方案
/*
   10110001 &  //初始值
   10110000    //减1后
  =10110000 &  //找出最低位的1
   10101111   //继续减1
  =10100000   //找出最低位的1
   ....       //直到找不出1即为0
*/

func BitCheck(i int)(count int){
	for i>0{
		count += 1
		i &= i-1
	}
	return
}

func main() {
	var a uint32
	a=0b00000000000100000000000000001011
	fmt.Println(hammingWeight(a))
}