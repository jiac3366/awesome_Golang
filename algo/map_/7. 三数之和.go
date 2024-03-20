package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			tmp := n1 + n2 + n3
			if tmp == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// 不等于（不重复）就不想移动了是吧？
				l++
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else if tmp > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

func main() {
	threeSum([]int{1, 2, 4, 3})
}
