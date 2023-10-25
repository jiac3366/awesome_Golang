package algo

// {1,2,2,4}
//        f
//  1 4
func RemoveElement(nums []int, val int) int {

	fast := 0
	slow := 0
	for _, item := range nums {
		if item != val {
			nums[slow] = nums[fast]
			fast++
			slow++
		} else {
			fast++
		}
	}
	return slow
}
