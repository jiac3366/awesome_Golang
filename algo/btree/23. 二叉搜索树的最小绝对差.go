package btree

import (
	"fmt"
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	var res []int
	traversal(root, &res)
	fmt.Println(res)
	return findMin(res)
}

func traversal(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	traversal(node.Left, arr)
	*arr = append(*arr, node.Val)
	traversal(node.Right, arr)
}

func findMin(arr []int) int {

	if len(arr) <= 1 {
		return 0
	}

	_min := math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		cur := arr[i+1] - arr[i]
		if cur < _min {
			_min = cur
		}
	}
	return _min
}

// ===========================================
func getMinimumDifference(root *TreeNode) int {
	var diff int
	diff = math.MaxInt
	var pre *TreeNode

	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Left)
		if pre != nil && root.Val-pre.Val < diff {
			diff = root.Val - pre.Val
		}
		pre = root
		traversal(root.Right)

	}
	traversal(root)
	return diff
}
