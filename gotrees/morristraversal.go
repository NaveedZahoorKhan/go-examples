package main

import "fmt"

/*
  Go program for
  Morris traversal for inorder
*/
// Binary Tree Node
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func getTreeNode(data int) *TreeNode {
	// return new TreeNode
	return &TreeNode{data, nil, nil}
}

type BinaryTree struct {
	root *TreeNode
}

func getBinaryTree() *BinaryTree {
	// return new BinaryTree
	return &BinaryTree{nil}
}

// Recursive function
// Display Inorder view of binary tree
func (this BinaryTree) inorder(node *TreeNode) {
	if node != nil {
		this.inorder(node.left)
		// Print node value
		fmt.Print("  ", node.data)
		this.inorder(node.right)
	}
}

// iterative inorder tree traversal
func (this BinaryTree) morrisInorder() {
	if this.root == nil {
		return
	}
	// Start to root node of tree
	var current *TreeNode = this.root
	var auxiliary *TreeNode = nil
	// iterating tree nodes
	for current != nil {
		if current.left == nil {
			// Print node value
			fmt.Print("  ", current.data)
			// When left child are empty then
			// visit to right child
			current = current.right
		} else {
			auxiliary = current.left
			for auxiliary.right != nil &&
				auxiliary.right != current {
				auxiliary = auxiliary.right
			}
			if auxiliary.right != current {
				// Change link
				auxiliary.right = current
				current = current.left
			} else {
				// Print node value
				fmt.Print("  ", current.data)
				// Unlink
				auxiliary.right = nil
				// Visit to right child
				current = current.right
			}
		}
	}
	fmt.Print("\n")
}

func main() {
	// Create new tree
	var tree *BinaryTree = getBinaryTree()
	// Create Binary Tree
	/*
	        4
	      /   \
	     8     3
	    / \   / \
	   1   6 7   2
	      /
	     9
	*/
	// Add tree node
	tree.root = getTreeNode(4)
	tree.root.left = getTreeNode(8)
	tree.root.left.left = getTreeNode(1)
	tree.root.right = getTreeNode(3)
	tree.root.right.right = getTreeNode(2)
	tree.root.right.left = getTreeNode(7)
	tree.root.left.right = getTreeNode(6)
	tree.root.left.right.left = getTreeNode(9)
	fmt.Println("\n Recursive Inorder")
	// Display inorder element
	tree.inorder(tree.root)
	fmt.Println("\n Morris Inorder")
	tree.morrisInorder()
}
