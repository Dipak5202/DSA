package avltree

type AvlNode struct {
	Data   int
	Left   *AvlNode
	Right  *AvlNode
	Height int
}

// rightRotate performs a right rotation on the given node.
func rightRotate(y *AvlNode) *AvlNode {
	x := y.Left
	T2 := x.Right

	// Perform rotation
	x.Right = y
	y.Left = T2

	// Update heights
	y.Height = max(getHeight(y.Left), getHeight(y.Right)) + 1
	x.Height = max(getHeight(x.Left), getHeight(x.Right)) + 1

	// Return new root
	return x
}

// leftRotate performs a left rotation on the given node.
func leftRotate(x *AvlNode) *AvlNode {
	y := x.Right
	T2 := y.Left

	// Perform rotation
	y.Left = x
	x.Right = T2

	// Update heights
	x.Height = max(getHeight(x.Left), getHeight(x.Right)) + 1
	y.Height = max(getHeight(y.Left), getHeight(y.Right)) + 1

	// Return new root
	return y
}

func getHeight(node *AvlNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func getBalance(node *AvlNode) int {
	if node == nil {
		return 0
	}
	return getHeight(node.Left) - getHeight(node.Right)
}

func InsertNode(root *AvlNode, key int) *AvlNode {
	if root == nil {
		return NewNode(key)
	}

	if key < root.Data {
		root.Left = InsertNode(root.Left, key)
	} else if key > root.Data {
		root.Right = InsertNode(root.Right, key)
	} else {
		return root // Duplicate keys are not allowed in BST
	}

	root.Height = 1 + max(getHeight(root.Left), getHeight(root.Right))

	balance := getBalance(root)

	if balance > 1 && key < root.Left.Data {
		return rightRotate(root)
	}

	if balance > 1 && key > root.Left.Data {
		root.Left = leftRotate(root.Left)
		return rightRotate(root)
	}

	if balance < -1 && key > root.Right.Data {
		return leftRotate(root)
	}

	if balance < -1 && key < root.Right.Data {
		root.Right = rightRotate(root.Right)
		return leftRotate(root)
	}

	return root
}

func NewNode(key int) *AvlNode {
	return &AvlNode{
		Data:   key,
		Left:   nil,
		Right:  nil,
		Height: 1,
	}
}

func InorderTraversal(root *AvlNode) {
	if root == nil {
		return
	}

	InorderTraversal(root.Left)
	print(root.Data, " ") // InorderTraversal(root.Right)
	InorderTraversal(root.Right)
}

func PreorderTraversal(root *AvlNode) {
	if root == nil {
		return
	}

	print(root.Data, " ") // PreorderTraversal(root.Left)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}

func PostorderTraversal(root *AvlNode) {
	if root == nil {
		return
	}

	PostorderTraversal(root.Left) // PostorderTraversal(root.Right)
	PostorderTraversal(root.Right)
	print(root.Data, " ")
}
