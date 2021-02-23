# 剑指 Offer 37. 序列化二叉树

## 题目

请实现两个函数，分别用来序列化和反序列化二叉树。

示例:

你可以将以下二叉树：

```go
    1
   / \
  2   3
     / \
    4   5
```

序列化为 `[1,2,3,null,null,4,5]`

## 思考

主要是使用BFS，用一个队列储存每一层的node

1. 初始化队列，包含root
2. 遍历队列，一边遍历，一边把子节点append进去
3. 直到末尾

在遍历的时候，你就可以进行序列化或反序列化操作了

1. 序列化：将当前node的值append到res数组里
2. 反序列化：从数组里取出值，新建node，当作其左节点和右节点

## 代码

1. 序列化

    ```go
    func serialize(root *TreeNode) []int {
        if root==nil{
            return nil
        }
        var (
            res  []int//结果数组
            tmp []*TreeNode//缓存每个节点的切片
        )
        tmp = append(tmp, root)
        for i := 0; i < len(tmp); i++ {
            if tmp[i]==nil{
                //如果节点为空，则结果用-999代替
                res = append(res, -999)
            }else{
                //否则，把当前值存入
                res = append(res, tmp[i].Val)
                //关键点！把当前节点的左右节点存起来，这样才能继续往下遍历
                tmp = append(tmp, tmp[i].Left)
                tmp = append(tmp, tmp[i].Right)
            }
        }
	    return res
    }

    ```

2. 反序列化

    ```go
    func deserialize(a []int)*TreeNode  {
        if len(a)==0{
            return nil
        }
        //同样需要缓存每个节点
        var tmp []*TreeNode
        root:=&TreeNode{
            Val: a[0],
        }
        tmp = append(tmp, root)
        //注意，这俩有2个游标，一个是值，一个是节点
        for i,j := 1,0; i < len(a); j++{
            //先取当前节点，开始是0，即root节点
            t:=tmp[j]
            //i初始值为1，即第二个节点
            //如果数组值为-999，说明当前root节点的左子树为空
            if a[i]!=-999{
                n:=new(TreeNode)
                n.Val=a[i]
                t.Left=n
                //同理，要将子节点入队
                tmp = append(tmp, n)
            }
            //注意！i要后移动一位
            i++
            if a[i]!=-999{
                n:=new(TreeNode)
                n.Val=a[i]
                t.Right=n
                tmp = append(tmp, n)
            }
            //i再向后移动一位，这时i已经移动了两位，对应节点t的两个子树的情况
            i++
        }
        return tmp[0]
    }
    ```
