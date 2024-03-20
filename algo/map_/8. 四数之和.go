package main

import "sort"

// 比三数之和多一层循环
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {

			if nums[i] > 0 && target <= 0 {
				continue
			}

			n1, n2 := nums[i], nums[j]
			l, r := j+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				tmp := n1 + n2 + n3 + n4
				if tmp == target {
					res = append(res, []int{i, j, l, r})
					for l < r && nums[l] == n3 {
						l++
					}
					for l < r && nums[r] == n4 {
						r--
					}
				} else if tmp < target {
					l++
				} else {
					r--
				}

			}

		}
	}
	return res
}
