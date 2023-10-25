package algo

import (
	"math"
)

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := float64(maxDepth(root.Left))
	right := float64(maxDepth(root.Right))

	return int(1 + math.Max(left, right))
}
