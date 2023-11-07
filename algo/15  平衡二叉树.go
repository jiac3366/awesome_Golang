package algo

import "math"

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1

}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)

	if leftHeight == -1 {
		return leftHeight
	}
	if rightHeight == -1 {
		return -1
	}

	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	} else {
		return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
	}
}
