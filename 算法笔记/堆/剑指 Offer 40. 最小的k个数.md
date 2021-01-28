# 剑指 Offer 40. 最小的k个数

## 题目

输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]

## 思考

1. **排序**后返回前k个数，使用快速排序进行。快速排序过程：
     1. 找基准值，一般取第一个为基准值
     2. 取第一个和最后一个为头指针和尾指针
     3. 遍历，条件是头小于尾
        1. 先从尾部向左，当有值小于基准值，则将此时的头部值换成尾部值，尾部值不变
        2. 再从头部往右，当有值大于基准值，则将此时的尾部值换成头部值，这样就把上一步的尾部值更新了，头部值等下一回和或者退出循环后更新
     4. 退出循环后，此时头部指向的值是基准值的位置，还没有更新，所以把基准值放进去
     5. 递归左半部分
     6. 递归右半部分
2. 使用**大顶堆**。大顶堆是一种数据结构，他的root节点比叶子节点大，即上面最小，下面最大.建堆后每次将堆顶元素和将要入堆的元素进行比较，如果大，直接抛弃，如果小，则将堆顶元素抛出，将此元素入堆，再进行堆化。这样遍历完所有数字后，大顶堆中的就是前k个最小值
   1. 建堆：找到最后一个节点的索引 `lastNode:=len(arr)-1`
   2. 找到其父节点的索引 `parent:=(lastNode-1)/2=(len(arr)/2-1)`
   3. 开始遍历进行堆化
      1. 找到此节点的左右两个孩子节点的索引：
         1. 左子节点：`left:=parent*2+1`
         2. 右子节点：`left:=parent*2+2`
      2. 比较3个节点值的大小关系，寻找最大值节点
      3. 交换父节点和最大值节点的值
      4. 因为最大值节点的值发生了改变，所以需要再往下进行堆化
   4. 堆化完成
   5. 遍历剩下的元素，将堆顶元素和将要入堆的元素进行比较，
      1. 如果大，直接抛弃，
      2. 如果小，则将堆顶元素抛出，将此元素入堆，再进行堆化。
   6. 结束

## 代码

1. 快速排序

    ```go
    //快速排序
    func getLeastNumbers(arr []int, k int) []int {
        if len(arr)==0||k>=len(arr){
            return arr
        }
        if k==0{
            return nil
        }
        quickSort(arr)
        fmt.Println(arr)
        return arr[:k]
    }

    func quickSort(arr []int) {
        if len(arr) <= 1 {
            return
        }
        p:=arr[0]
        low,big:=0,len(arr)-1
        for low<big{
            for low<big&&arr[big]>=p {
                big--
            }
            arr[low]=arr[big]
            for low<big&&arr[low]<=p {
                low++
            }
            arr[big]=arr[low]
        }
        arr[low]=p
        quickSort(arr[:low])
        quickSort(arr[low+1:])
    }
    ```

2. 大顶堆

    ```go
    func getLeastNumbersII(arr []int, k int) []int {
        if len(arr)==0||k>=len(arr){
            return arr
        }
        if k==0{
            return nil
        }
        heap:=arr[:k]
        fmt.Println(heap)
        makeHeap(heap)
        fmt.Println(heap)

        for i := k; i <len(arr); i++ {
            if arr[i]<heap[0]{
                heap[0]=arr[i]
                heapify(heap,0)
            }
        }
        return heap
    }

    func makeHeap(arr []int)  {
        l:= len(arr)
        lastNode:=l-1
        lastParent:=(lastNode-1)/2
        for i := lastParent; i >=0 ; i-- {
            heapify(arr,i)
        }
    }

    func heapify(arr []int,i int)  {
        left:=i*2+1
        right:=left+1
        max:=i
        if left<len(arr)&&arr[left]>arr[max]{
            max=left
        }
        if  right<len(arr)&&arr[right]>arr[max]{
            max=right
        }

        if max!=i{
            arr[i],arr[max]=arr[max],arr[i]
            heapify(arr,max)
        }
    }

    ```
