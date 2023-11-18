package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := float64(maxDepth(root.Left))
	right := float64(maxDepth(root.Right))

	return int(1 + math.Max(left, right))
}

// 前序遍历
// 暂时注释
//var result int
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	_min = math.MaxInt64
	_max = math.MinInt64
)

func isValidBST(root *TreeNode) bool {

	findMax(root.Left)

	if root.Val <= _max {
		return false
	}

	findMin(root.Right)
	if root.Val >= _min {
		return false
	}

	isLeftValid := isValidBST(root.Left)
	isRightValid := isValidBST(root.Right)

	return isLeftValid && isRightValid

}

func findMaxValueRecursive(root *TreeNode) int {
	if root == nil {
		return math.MinInt64
	}

	// 递归计算左右子树的最大值
	leftMax := findMaxValueRecursive(root.Left)
	rightMax := findMaxValueRecursive(root.Right)

	// 计算当前节点和左右子树的最大值
	currentMax := int(math.Max(float64(root.Val), math.Max(float64(leftMax), float64(rightMax))))

	return currentMax
}

func findMinValueRecursive(root *TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}

	// 递归计算左右子树的最大值
	leftMax := findMinValueRecursive(root.Left)
	rightMax := findMinValueRecursive(root.Right)

	// 计算当前节点和左右子树的最大值
	currentMax := int(math.Min(float64(root.Val), math.Min(float64(leftMax), float64(rightMax))))

	return currentMax
}
