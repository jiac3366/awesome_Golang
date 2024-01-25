package btree

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > high {
		// 向左遍历返回 val <= high的第一个节点
	}

	if root.Val < low {
		// 向右遍历 返回 val >= low 的第一个节点
	}

	// 其实上述代码已经 ok 了，记住，删除搜索树节点就是保留左子树，返回右子树
}
