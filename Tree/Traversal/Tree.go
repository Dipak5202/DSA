package traversal

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data, Left: nil, Right: nil}
}

func InorderTraversal(root *Node) {
	if root == nil {
		return
	}

	InorderTraversal(root.Left)
	print(root.Data, " ")
	InorderTraversal(root.Right)
}

func PreorderTraversal(root *Node) {
	if root == nil {
		return
	}

	print(root.Data, " ") // PreorderTraversal(root.Left)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}

func PostorderTraversal(root *Node) {
	if root == nil {
		return
	}

	PostorderTraversal(root.Left) // PostorderTraversal(root.Right)
	PostorderTraversal(root.Right)
	print(root.Data, " ")
}

// in which every parent node/internal node has either two or no children.
func FullBinaryTree(root *Node) bool {

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		return FullBinaryTree(root.Left) && FullBinaryTree(root.Right)
	}
	return false
}

func Depth(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := Depth(root.Left)
	rightDepth := Depth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func IsPerfectBinaryTree(root *Node, depth int, level int) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return depth == level+1
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	return IsPerfectBinaryTree(root.Left, depth, level+1) && IsPerfectBinaryTree(root.Right, depth, level+1)
}

func Hight(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := Hight(root.Left)
	rightHeight := Hight(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
