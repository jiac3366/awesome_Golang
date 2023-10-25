package algo

// [-4,-1,0,3,10]
//   0  1 9 16  100

func SortedSquares(nums []int) []int {
	newNum := make([]int, len(nums))

	lo, hi, idx := 0, len(nums)-1, len(nums)-1

	for lo <= hi {
		loM, hiM := nums[lo]*nums[lo], nums[hi]*nums[hi]
		if loM > hiM {
			newNum[idx] = loM
			lo++
		} else {
			newNum[idx] = hiM
			hi--
		}
		idx--
	}

	return newNum

}
