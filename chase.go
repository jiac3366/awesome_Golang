package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6}
	slice1 := arr[1:4:5] //{low:high:max} 最多再扩张一个元素
	//max超出 len(arr)
	//slice2 := arr[1:4:7] //panic
	fmt.Println(slice1)     //[1,2,3]
	slice3 := slice1[1:3:4] //[2,3] 大于4会panic
	fmt.Println(slice3)
	slice4 := slice1[1:4] //[2,3] 大于4会panic
	fmt.Println(slice4)   //[2 3 4] 出现slice1中没有的元素
}
