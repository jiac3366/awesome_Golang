package btree

import (
	"container/list"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)

	level := 0

	for queue.Len() > 0 {
		size := queue.Len()
		flag := false
		for ; size > 0; size-- {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			// 找出 没有左子树和右子树的节点，提前退出
			if node.Left == nil && node.Right == nil {
				flag = true
				break
			}
		}
		level += 1
		if flag {
			break
		}
	}
	return level

}
