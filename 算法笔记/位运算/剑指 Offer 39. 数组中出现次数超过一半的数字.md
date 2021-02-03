# 剑指 Offer 39. 数组中出现次数超过一半的数字

## 题目

数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例1:
输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]  
输出: 2

## 思考

1. 排序，然后返回 `n/2` 处的元素即可
2. 哈希表，遍历数组，对应的value加一，然后遍历哈希表，找到value大于 n/2 的key即可
3. 摩尔投票法，选一个基准值，票数为1，然后遍历，和基准值相同，则加一票；不同，则减一票；为0时，放弃这个数，换成当前数字的下一个数字。
4. **位运算**：对于32位数字的每一位，遍历数组，统计如果当前位为1的个数；如果大于一半，则说明众数的这一位一定是1，否则就是0；这样遍历完32位后，就能得到众数的每一位的情况，进而得到众数。

## 代码

1. 位运算

    ```go
    func majorityElement(nums []int) int {
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

    ```

2. 摩尔投票法

    ```go
    func majorityElement3(nums []int) int {
        l:=len(nums)
        if l==0{
            return 0
        }
        if l<=2{
            return nums[0]
        }
        res,count:=1,1
        for i:=1;i<l;i++{
            if count==0{
                res=nums[i]
                continue
            }
            if nums[i]==res{
                count++
            }else{
                count--
            }
        }
        return res
    }
    ```
