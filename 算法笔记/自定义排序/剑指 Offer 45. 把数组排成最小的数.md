# 剑指 Offer 45. 把数组排成最小的数

## 题目

输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

示例 1:
输入: [10,2]
输出: "102"

示例 2:
输入: [3,30,34,5,9]
输出: "3033459"

## 思考

我的思考：

1. 穷举，然后比较🐷
2. 把数组按数字的位数分组，然后在每组进行排序。。。然后没了🐷

解题答案：
先全部转换为字符串，然后进行自定义快速排序，即 如果 `x+y < y+x`,则 `x < y`。这样最终排完序之后，字符串数组拼接后就是最小的 🧠
借着这题，把几种排序再复习一下

时间复杂度：
![sjfzd](https://www.geekxh.com/assets/img/sort-1.5b9670d7.png)

1. 冒泡法
两两比较相邻记录的关键字，如果反序则交换，直到没有反序的记录为止
直接看图：
![mpfpx](https://www.geekxh.com/assets/img/bubbleSort.b7d216a5.gif)
最好：正序 `log(N)`,最坏：倒序 `log(N^2)`
2. 基数排序
根据各个数位数的大小，从个位一直遍历到最高位，每次遍历时放到对应位的桶上，然后再顺序遍历桶收回，就实现了按照这一位的大小进行排序；把所有位数遍历完后，结果就是排好序了。
![jspx](https://www.geekxh.com/assets/img/radixSort.6690b105.gif)
时间：`log(n*k)`,空间`log(k*n)`
3. 计数排序
使用map进行缓存，一个key的value要++,然后遍历map，每个元素对应的位置就是小于他的元素的value的和
![jspx](https://www.geekxh.com/assets/img/countingSort.827d96b8.gif)
4. 选择排序
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
重复第二步，直到所有元素均排序完毕。
![xzpx](https://www.geekxh.com/assets/img/selectionSort.44be35da.gif)
5. 插入排序
选定第一个为基准，然后将遍历后面的n-1个数，根据大小放到基准值的左右
![crpx](https://www.geekxh.com/assets/img/insertionSort.be81c151.gif)
6. 希尔排序(分组间隔插入)
是插入排序的一种优化，将数组按照不同的间隔分为几组，然后对这几组进行插入排序；每次排序当间隔为1时，相当于对整对进行插入排序，排完就结束了。
7. 归并排序(分治，双指针，递归)
从上到下递归，每次递归时：
   1. 如果传入的数组长度为1或0，返回原数组
   2. 否则，继续往下分为2半
   3. 对这2半再进行递归调用本函数
   4. 将递归的返回结果进行合并，合并时使用双指针法，小的上去，指针前进
![gbpx](https://www.geekxh.com/assets/img/mergeSort.9541d116.gif)
8. 快速排序(双指针，递归)
   1. 找基准值(a[0])
   2. 双指针，先尾部向左，比基准值小就和头指针交换位置
   3. 然后头部向右，比基准值大就和尾指针交换位置
   4. 重复2、3两步，直到头尾指针相遇
   5. 此时a[0]已经到了他合适的位置了，接下来递归左半部分和右半部分，直到数组长度小于2
![kspx](https://www.geekxh.com/assets/img/quickSort.71c0f1c0.gif)
时间：最坏O(n2),平均O(n*logn)
9. 堆排序
利用大顶堆或者小顶堆进行排序。大小顶堆的原理和实现见[最小的k个数](../堆/剑指%20Offer%2040.%20最小的k个数.md)
   1. 将待排序序列构建成一个堆 H[0……n-1]，根据（升序降序需求）选择大顶堆或小顶堆；
   2. 把堆首（最大值）和堆尾互换；
   3. 把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
   4. 重复步骤 2，直到堆的尺寸为 1。
![dpx](https://www.geekxh.com/assets/img/heapSort.658d0f58.gif)
代码见[堆排序](../../DataStructure/0911堆排序.go)

## 代码

```go
func minNumber(nums []int) string {
	l:= len(nums)
	if l==0{
		return ""
	}
	tmp:=make([]string,l)
	for i := 0; i < l; i++ {
		tmp[i]=strconv.Itoa(nums[i])
	}
	tmp=quickSort(tmp)
	res:=new( strings.Builder)
	for i := 0; i < l; i++ {
		res.WriteString(tmp[i])
	}
	return res.String()
}

func quickSort(strs []string) []string {
	l:= len(strs)
	if l<2{
		return strs
	}
	base:=strs[0]
	left,right:=0,l-1

	for left < right {
		for right>left&&strs[right]+base>=base+strs[right]{
			right--
		}
		strs[right],strs[left]=strs[left],strs[right]
		for right>left&&strs[left]+base<=base+strs[right]{
			left++
		}
		strs[right],strs[left]=strs[left],strs[right]
	}
	quickSort(strs[:left])
	quickSort(strs[left+1:])
	return strs
}

```
