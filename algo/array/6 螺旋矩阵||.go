package main

import "fmt"

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	up, down, left, right := 0, n-1, 0, n-1
	cur := 1
	for cur <= n*n {
		// up: left --> right
		for i := left; i <= right; i++ {
			ans[up][i] = cur
			cur++
			fmt.Println("1")
		}
		up++

		// right: up --> down
		for i := up; i <= down; i++ {
			ans[i][right] = cur
			cur++
		}
		right--

		// down: right --> left
		for i := right; i >= left; i-- {
			ans[down][i] = cur
			cur++
		}
		down--

		// left: down --> up
		for i := down; i >= up; i-- {
			ans[i][left] = cur
			cur++
		}
		left++

	}

	return ans
}
