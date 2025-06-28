package problems

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}

	// Check left and right subtrees
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}
