/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package btree

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Val <= findMax(root.Left) {
		return false
	}

	if root.Val >= findMin(root.Right) {
		return false
	}

	isLeftValid := isValidBST(root.Left)
	isRightValid := isValidBST(root.Right)

	return isLeftValid && isRightValid

}

func findMax(root *TreeNode) int {
	if root == nil {
		return math.MinInt64
	}

	// 递归计算左右子树的最大值
	leftMax := findMax(root.Left)
	rightMax := findMax(root.Right)

	// 计算当前节点和左右子树的最大值
	currentMax := int(math.Max(float64(root.Val), math.Max(float64(leftMax), float64(rightMax))))

	return currentMax
}

func findMin(root *TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}

	// 递归计算左右子树的最大值
	leftMax := findMin(root.Left)
	rightMax := findMin(root.Right)

	// 计算当前节点和左右子树的最大值
	currentMax := int(math.Min(float64(root.Val), math.Min(float64(leftMax), float64(rightMax))))

	return currentMax
}
