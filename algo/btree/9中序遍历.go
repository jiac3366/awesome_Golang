package btree

import "container/list"

func inorderTraversal(root *TreeNode) []int {

	ans := []int{}

	if root == nil {
		return ans
	}

	st := list.New()
	cur := root

	for st.Len() > 0 || cur != nil {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			cur = st.Remove(st.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}

	return ans

}
