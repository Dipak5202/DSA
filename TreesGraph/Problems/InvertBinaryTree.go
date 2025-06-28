package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap left and right
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)

	return root
}
