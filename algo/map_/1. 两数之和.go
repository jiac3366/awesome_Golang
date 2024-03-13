package main

func twoSum(nums []int, target int) []int {
	revert := make(map[int]int)
	for index, num := range nums {
		valIndex, ok := revert[target-num]
		if ok {
			return []int{index, valIndex}
		} else {
			revert[num] = index
		}
	}
	return []int{}
}
