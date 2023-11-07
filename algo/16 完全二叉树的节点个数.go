package algo

import "container/list"

func countNodes(root *TreeNode) int {
	var num int

	if root == nil {
		return num
	}

	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		num += 1
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}
	return num
}
