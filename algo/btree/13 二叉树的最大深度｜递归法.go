package btree

import (
	"math"
)

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

func maxDepth(root *TreeNode) int {
	result = 0
	if root == nil {
		return result
	}
	_maxDepth(root, 1)
	return result
}

func _maxDepth(node *TreeNode, level int) {
	result = int(math.Max(float64(level), float64(result)))

	if node.Left == nil && node.Right == nil {
		return
	}

	if node.Left != nil {
		_maxDepth(node.Left, level+1)
	}

	if node.Right != nil {
		_maxDepth(node.Right, level+1)
	}
}
