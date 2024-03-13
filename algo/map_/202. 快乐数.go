package main

import "strconv"

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	res := make(map[int]int)
	for {
		multi := 0
		for _, s := range strconv.Itoa(n) {
			i, _ := strconv.Atoi(string(s))
			multi += i * i
		}
		if multi == 1 {
			return true
		}
		_, ok := res[n]
		if !ok {
			res[n] = multi
			n = multi
		} else {
			return false
		}
	}

}
