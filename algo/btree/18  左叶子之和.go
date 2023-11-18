package btree

import (
	"container/list"
)

func findBottomLeftValue(root *TreeNode) int {

	st := list.New()
	st.PushBack(root)

	var lastNode int
	for st.Len() > 0 {
		l := st.Len()
		for i := 0; i < l; i++ {
			node := st.Remove(st.Front()).(*TreeNode)
			if i == l-1 {
				lastNode = node.Val
			}
			if node.Right != nil {
				st.PushBack(node.Right)
			}

			if node.Left != nil {
				st.PushBack(node.Left)
			}
		}
	}
	return lastNode
}
