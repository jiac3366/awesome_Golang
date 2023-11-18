package array

func Search(nums []int, target int) int {
	_len := len(nums)
	if _len == 0 {
		return -1
	}
	lo := 0
	hi := _len - 1
	for lo <= hi {
		mid := (lo + hi) / 2
		if target > nums[mid] {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1

}
