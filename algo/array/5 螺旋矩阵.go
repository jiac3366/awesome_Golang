package array

// abca
func LengthOfLongestSubstring(s string) int {
	counter := make(map[byte]int)
	left, right := 0, 0

	maxLen := 0
	for right < len(s) {
		newLetter := s[right]
		counter[newLetter]++
		right++
		for counter[newLetter] > 1 {
			leftLetter := s[left]
			counter[leftLetter]--

			left++
		}
		curLen := right - left
		if curLen > maxLen {
			maxLen = curLen
		}

	}

	return maxLen
}
