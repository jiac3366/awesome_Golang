package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//[5,3,6,2,4,null,7]

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key && root.Left == nil && root.Right == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
	}

	if root.Val == key {
		// 把当前要删除节点的左子树 作为 要删除节点的右子树的最左节点 的左子树
		tmp := root.Right
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		root = tmp
		return root
	}

	if root.Val < key {
		root.Right = traversal(root.Right, key)
	}

	if root.Val > key {
		root.Left = traversal(root.Left, key)
	}
	return root
}
