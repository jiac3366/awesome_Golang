package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	var traversal func(root *TreeNode, val int) bool
	traversal = func(root *TreeNode, val int) bool {
		if root == nil {
			return true
		}

		if root.Val < val {
			if traversal(root.Right, val) {
				root.Left = &TreeNode{Val: val}
			}
		} else {
			if traversal(root.Right, val) {
				root.Right = &TreeNode{Val: val}
			}
		}

		return false

	}
	traversal(root, val)
	return root

}
