package btree

import (
	"container/list"
)

func sumOfLeftLeaves(root *TreeNode) int {

	st := list.New()
	st.PushBack(root)

	var sum int
	for st.Len() > 0 {
		l := st.Len()
		for i := 0; i < l; i++ {
			node := st.Remove(st.Front()).(*TreeNode)

			if node.Right != nil {
				st.PushBack(node.Right)
			}

			if node.Left != nil {
				st.PushBack(node.Left)
				if node.Left.Left == nil && node.Left.Right == nil {
					sum += node.Left.Val
				}
			}
		}
	}
	return sum
}
