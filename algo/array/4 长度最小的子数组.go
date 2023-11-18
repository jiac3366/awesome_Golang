package array

func MinSubArrayLen(target int, nums []int) int {
	// [2,3,1,2,4,3]  7
	cur_num := 0
	start := 0
	l := len(nums)
	result_len := l + 1

	for i := 0; i < l; i++ {
		cur_num += nums[i]
		for cur_num >= target {
			sub_len := i - start + 1
			if sub_len < result_len {
				result_len = sub_len
			}
			cur_num -= nums[start]
			start++
		}
	}

	return result_len
}
