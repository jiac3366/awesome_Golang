package algo

import "strconv"

// 递归法
func binaryTreePaths1(root *TreeNode) []string {
	res := []string{}

	var traver func(node *TreeNode, s string)
	traver = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			s += strconv.Itoa(node.Val)
			res = append(res, s)
			return
		}

		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			traver(node.Left, s)
		}
		if node.Right != nil {
			traver(node.Right, s)
		}

	}
	traver(root, "")
	return res

}

func binaryTreePaths(root *TreeNode) []string {
	///
}
