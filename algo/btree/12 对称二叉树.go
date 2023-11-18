package btree

import (
	"container/list"
)

func isSymmetric(root *TreeNode) bool {
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(left *TreeNode, right *TreeNode) bool {
	if left != nil && right == nil {
		return false
	}
	if left == nil && right != nil {
		return false
	}
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil {
		// 不能简单地理解为 left right 是相等的就是对称的,还需要判断其子树
		//if left.Val == right.Val {
		//	return true
		//}

		if left.Val != right.Val {
			return false
		}
	}

	isInsideEqual := isSymmetricTree(left.Right, right.Left)
	isOutsideEqual := isSymmetricTree(left.Left, right.Right)
	return isOutsideEqual && isInsideEqual
}

// todo
func isSymmetric(root *TreeNode) bool {
	st := list.New()
	if root != nil {
		st.PushBack(root.Right)
		st.PushBack(root.Left)
	}

	for st.Len() > 0 {
		left := st.Remove(st.Back()).(*TreeNode)
		right := st.Remove(st.Back()).(*TreeNode)

	}

}
