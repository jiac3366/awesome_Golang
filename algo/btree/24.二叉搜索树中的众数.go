package btree

func findMode(root *TreeNode) []int {

	var pre *TreeNode
	var MaxCount = 1
	var count = 1
	var res []int
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {

		if node == nil {
			return
		}
		traversal(node.Left)

		if pre == nil {
			count = 1
		} else if pre.Val != node.Val {
			count = 1
		} else {
			count += 1

		}

		if count == MaxCount {
			res = append(res, node.Val)
		}

		if count > MaxCount {
			res = []int{node.Val}
			MaxCount = count
		}

		pre = node
		traversal(node.Right)

	}
	traversal(root)
	return res

}
