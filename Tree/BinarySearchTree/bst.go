package binarysearchtree

import traversal "tree/Traversal"

func Insert(node *traversal.Node, key int) *traversal.Node {

	if node == nil {
		return traversal.NewNode(key)
	}

	if key < node.Data {
		node.Left = Insert(node.Left, key)
	} else {
		node.Right = Insert(node.Right, key)
	}
	return node
}

func InsertUsingFor(node *traversal.Node, key int) *traversal.Node {
	if node == nil {
		return traversal.NewNode(key)
	}

	current := node // Use a temporary pointer for traversal

	for {
		if key < current.Data {
			if current.Left == nil {
				current.Left = traversal.NewNode(key)
				break
			}
			current = current.Left
		} else if key > current.Data {
			if current.Right == nil {
				current.Right = traversal.NewNode(key)
				break
			}
			current = current.Right
		} else {
			break // Duplicate keys are not allowed in BST
		}
	}
	return node // Return the original root node
}

func DeleteNode(root *traversal.Node, key int) *traversal.Node {
	if root == nil {
		return root
	}

	if key < root.Data {
		root.Left = DeleteNode(root.Left, key)
	} else if key > root.Data {
		root.Right = DeleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minNode := FindMin(root.Right)
		root.Data = minNode.Data
		root.Right = DeleteNode(root.Right, minNode.Data)
	}
	return root
}

func FindMin(node *traversal.Node) *traversal.Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func Search(node *traversal.Node, key int) *traversal.Node {
	if node == nil || node.Data == key {
		return node
	}

	if key < node.Data {
		return Search(node.Left, key)
	}
	return Search(node.Right, key)
}
