package main

import (
	avltree "tree/AVLTree"
	binarysearchtree "tree/BinarySearchTree"
	traversal "tree/Traversal"
)

func main() {

	// Traversal()

	// FullBinaryTree()

	// Depth()

	// IsPerfectBinaryTree()

	// Binarysearchtree()

	// BinarysearchtreeUsingFor()
	AVLtree()
}

func BinarysearchtreeUsingFor() {

	println("Binary Search Tree using for loop")
	root := traversal.NewNode(7)
	root = binarysearchtree.InsertUsingFor(root, 2)
	root = binarysearchtree.InsertUsingFor(root, 1)
	root = binarysearchtree.InsertUsingFor(root, 3)
	root = binarysearchtree.InsertUsingFor(root, 6)
	root = binarysearchtree.InsertUsingFor(root, 5)
	root = binarysearchtree.InsertUsingFor(root, 9)
	root = binarysearchtree.InsertUsingFor(root, 8)
	root = binarysearchtree.InsertUsingFor(root, 10)
	root = binarysearchtree.InsertUsingFor(root, 4)
	root = binarysearchtree.InsertUsingFor(root, 11)

	println("Inorder Traversal of the BST:")
	traversal.InorderTraversal(root)
	println()

	traversal.PreorderTraversal(root)
	println()

	traversal.PostorderTraversal(root)
	println()
}
func Binarysearchtree() {
	root := traversal.NewNode(7)
	root = binarysearchtree.Insert(root, 2)
	root = binarysearchtree.Insert(root, 1)
	root = binarysearchtree.Insert(root, 3)
	root = binarysearchtree.Insert(root, 6)
	root = binarysearchtree.Insert(root, 5)
	root = binarysearchtree.Insert(root, 9)
	root = binarysearchtree.Insert(root, 8)
	root = binarysearchtree.Insert(root, 10)
	root = binarysearchtree.Insert(root, 4)
	root = binarysearchtree.Insert(root, 11)

	println("Inorder Traversal of the BST:")
	traversal.InorderTraversal(root)
	println()

	traversal.PreorderTraversal(root)
	println()

	traversal.PostorderTraversal(root)
	println()
}
func IsPerfectBinaryTree() {
	root := traversal.NewNode(1)
	root.Left = traversal.NewNode(2)
	root.Right = traversal.NewNode(3)
	root.Left.Left = traversal.NewNode(4)
	root.Left.Right = traversal.NewNode(5)
	// root.Right.Left = traversal.NewNode(6)
	root.Right.Right = traversal.NewNode(7)

	if traversal.IsPerfectBinaryTree(root, traversal.Depth(root), 0) {
		println("The tree is a perfect binary tree.")
	} else {
		println("The tree is not a perfect binary tree.")
	}
}
func Depth() {
	root := traversal.NewNode(1)
	root.Left = traversal.NewNode(2)
	root.Right = traversal.NewNode(3)
	root.Left.Left = traversal.NewNode(4)
	root.Left.Right = traversal.NewNode(5)
	root.Right.Left = traversal.NewNode(6)
	root.Right.Right = traversal.NewNode(7)
	root.Left.Right.Left = traversal.NewNode(8)
	println("Depth of the tree is:", traversal.Depth(root))
}
func FullBinaryTree() {
	root := traversal.NewNode(1)
	root.Left = traversal.NewNode(2)
	root.Right = traversal.NewNode(3)
	root.Left.Left = traversal.NewNode(4)
	// root.Left.Right = traversal.NewNode(5)
	root.Right.Left = traversal.NewNode(6)
	root.Right.Right = traversal.NewNode(7)

	if traversal.FullBinaryTree(root) {
		println("The tree is a full binary tree.")
	} else {
		println("The tree is not a full binary tree.")
	}
}
func Traversal() {

	root := traversal.NewNode(1)
	root.Left = traversal.NewNode(2)
	root.Right = traversal.NewNode(3)
	root.Left.Left = traversal.NewNode(4)
	root.Left.Right = traversal.NewNode(5)
	root.Right.Left = traversal.NewNode(6)
	root.Right.Right = traversal.NewNode(7)
	println("Inorder Traversal:")
	traversal.InorderTraversal(root)
	println()

	println("Preorder Traversal:")
	traversal.PreorderTraversal(root)
	println()

	println("Postorder Traversal:")
	traversal.PostorderTraversal(root)
	println()
}

func AVLtree() {
	root := avltree.NewNode(1)
	root = avltree.InsertNode(root, 2)
	root = avltree.InsertNode(root, 3)
	root = avltree.InsertNode(root, 4)
	root = avltree.InsertNode(root, 5)
	root = avltree.InsertNode(root, 6)
	root = avltree.InsertNode(root, 7)
	root = avltree.InsertNode(root, 8)
	root = avltree.InsertNode(root, 9)

	println("Inorder Traversal of the AVL Tree:")
	avltree.InorderTraversal(root)
	println()

	avltree.PreorderTraversal(root)
	println()

	avltree.PostorderTraversal(root)
	println()
}
