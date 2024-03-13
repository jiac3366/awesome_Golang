package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	const lLen = 1000
	both := make([]int, lLen)
	for i := 0; i <= lLen-1; i++ {
		both[i] = -1
	}

	res := make([]int, 0)
	for _, item := range nums1 {
		both[item] = 1
	}

	for _, item := range nums2 {
		if both[item] == 1 {
			res = append(res, item)
		}
	}

	return res
}

func main() {
	fmt.Printf("%+v", intersection([]int{1, 2, 3}, []int{2, 3, 4}))
}
