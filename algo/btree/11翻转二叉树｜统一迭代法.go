package btree

import (
	"container/list"
)

func invertTree(root *TreeNode) *TreeNode {

	// 统一迭代法

	if root == nil {
		return root
	}
	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		e := st.Back()
		st.Remove(e) // e有可能是个标志位，其值为 nil,没办法做 类型转换.(*TreeNode)
		if e.Value != nil {
			node := e.Value.(*TreeNode)
			if node.Right != nil {
				st.PushBack(node.Right)
			}
			if node.Left != nil {
				st.PushBack(node.Left)
			}
			st.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符，表示它的左右子节点已经入栈，下一次访问到这个节点，直接访问就可以了
			st.PushBack(nil)
		} else {
			// 处理
			midNode := st.Remove(st.Back()).(*TreeNode)
			midNode.Left, midNode.Right = midNode.Right, midNode.Left
		}
	}
	return root

}
