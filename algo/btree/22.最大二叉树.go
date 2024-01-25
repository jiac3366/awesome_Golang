package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := findMaxIndex(nums)
	root := &TreeNode{Val: nums[index]}
	leftNUms := nums[:index]
	rightNums := nums[index+1:]

	root.Left = constructMaximumBinaryTree(leftNUms)
	root.Right = constructMaximumBinaryTree(rightNums)
	return root
}

func findMaxIndex(nums []int) int {

	best := 0
	max := nums[0]
	for index, item := range nums {
		if item > max {
			max = item
			best = index
		}
	}
	return best
}
