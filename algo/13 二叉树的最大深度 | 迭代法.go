package algo

import (
	"container/list"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)

	level := 0

	for queue.Len() > 0 {
		size := queue.Len()
		for ; size > 0; size-- {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		level += 1
	}
	return level

}
