package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	cache := make(map[int]int)
	for _, val1 := range nums1 {
		for _, val2 := range nums2 {
			cache[val1+val2] += 1
		}
	}

	res := 0
	for _, val3 := range nums3 {
		for _, val4 := range nums4 {
			if val, ok := cache[0-(val3+val4)]; ok {
				res += val
			}
		}

	}
	return res
}
