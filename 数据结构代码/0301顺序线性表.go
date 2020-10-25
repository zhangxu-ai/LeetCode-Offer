package main

import (
    "errors"
)

type SqList struct {
    data []interface{}
    length int
}

func (sl *SqList) Get(index int) (interface{},error){
    if index >=sl.length{
        return -1,errors.New("out of index")
    }else if index <1{
        return -1,errors.New("invalid index")
    }
    return sl.data[index-1],nil
}

func (sl *SqList) InsertOne(e interface{}, p int) bool {
    if p <1|| p >sl.length+1{
        return false
    }
    if p <=sl.length{
        for  i:=sl.length;i>= p;i--{
            if i==sl.length{
                sl.data=append(sl.data,sl.data[i-1])
            }
            sl.data[i]=sl.data[i-1]
        }
    }
    sl.data[p-1]=e
    return true
}

func (sl *SqList) DeleteOne(p int) (interface{},error) {
    if p <1|| p >sl.length{
        return nil,errors.New("invalid position")
    }
    e:=sl.data[p-1]
    //模拟删除操作
    sl.data[p-1]=nil
    //前移
    if p<sl.length{
        for i:= p-1;i<sl.length-1;i++{
            sl.data[i]=sl.data[i+1]
        }
    }
    sl.data=sl.data[:sl.length-1]
    sl.length--
    return e,nil
}

//func main(){
//    sq:=&SqList{
//        data:   []interface{}{1,2,3,4,5},
//        length: 5,
//    }
//    fmt.Println(sq.Get(1))
//    if sq.InsertOne(3.4,4){
//        fmt.Println("insert ok")
//    }
//}
