package main

//折半查找
func BinarySearch(a []int,key int) int {
	low:=0
	high:=len(a)-1
	mid:=(low+high)/2+1
	for low<=high{
		if a[mid]<key{
			low=mid+1
		}else if a[mid]>key{
			high=mid-1
		}else {
			return mid
		}
	}
	return -1
}
