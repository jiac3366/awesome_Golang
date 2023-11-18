package btree

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return traversal(root, targetSum, 0)
}

func traversal(node *TreeNode, targetSum int, levelSum int) bool {

	if node.Left == nil && node.Right == nil {
		if levelSum+node.Val == targetSum {
			return true
		} else {
			return false
		}
	}

	leftRes, rightRes := false, false
	levelSum += node.Val
	if node.Left != nil {
		leftRes = traversal(node.Left, targetSum, levelSum)
	}

	if node.Right != nil {
		rightRes = traversal(node.Right, targetSum, levelSum)
	}

	return leftRes || rightRes

}
