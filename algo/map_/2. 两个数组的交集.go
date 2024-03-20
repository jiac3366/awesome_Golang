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

func intersection2(nums1 []int, nums2 []int) []int {
	record := make(map[int]struct{}, 0)

	for _, num := range nums1 {
		record[num] = struct{}{}
	}

	res := make([]int, 0)
	for _, num := range nums2 {
		if _, ok := record[num]; ok {
			res = append(res, num)
		}
		// 这里要及时删掉，因为 nums2 不保证重复，我们有可能会给 res 加入重复的元素
		delete(record, num)
	}
	return res
}

func main() {
	fmt.Printf("%+v", intersection([]int{1, 2, 3}, []int{2, 3, 4}))
}
