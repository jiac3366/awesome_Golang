package btree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil // 因为对于inorder =[2,1]; postorder =[2,1] 这些情况，传入右分支的inorder和postorder都是空，所以既然postorder为空，直接返回 nil 即可
	}

	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}

	hash := make(map[int]int)
	for i, v := range inorder { // 用map保存中序序列的数值对应位置
		hash[v] = i
	}
	rootV := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootV}
	InIdx := hash[rootV]

	leftInOrder := inorder[:InIdx]
	rightInOrder := inorder[InIdx+1:]

	leftPostOrder := postorder[:len(leftInOrder)]
	rightPostOrder := postorder[len(leftInOrder) : len(postorder)-1]

	root.Left = buildTree(leftInOrder, leftPostOrder)
	root.Right = buildTree(rightInOrder, rightPostOrder)
	return root
}
