package btree

import "container/list"

func levelOrder(root *TreeNode) [][]int {

	ans := [][]int{}
	if root == nil {
		return ans
	}
	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		tmp := []int{}
		for i := 0; i < st.Len(); i++ {
			node := st.Remove(st.Front()).(*TreeNode)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				st.PushBack(node.Left)
			}
			if node.Right != nil {
				st.PushBack(node.Right)
			}
		}
		ans = append(ans, tmp)

	}
	return ans
}
