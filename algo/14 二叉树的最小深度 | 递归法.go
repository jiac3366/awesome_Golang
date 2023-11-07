package algo

import "math"

// 后序遍历
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	}

	return 1 + int(math.Min(
		float64(minDepth(root.Left)),
		float64(minDepth(root.Right)),
	))
}

// 前序遍历
var result float64

func minDepth(root *TreeNode) int {
	result = math.Inf(1)
	if root == nil {
		return 0
	}
	_minDepth(root, 1)
	return int(result)
}

func _minDepth(node *TreeNode, level int) {

	if node.Left == nil && node.Right == nil {
		result = math.Min(float64(level), float64(result))
		return
	}

	if node.Left != nil {
		_minDepth(node.Left, level+1)
	}

	if node.Right != nil {
		_minDepth(node.Right, level+1)
	}
}
