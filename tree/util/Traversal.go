package util

import (
	"fmt"
)

// InorderRecursive : Inorder recursive traversal
func InorderRecursive(root *Node) {
	fmt.Println("as", root)
	if root == nil {
		InorderRecursive(root.left)
		fmt.Printf("%d ", root.data)
		InorderRecursive(root.right)
	}
}

// Inorder : traverses the tree
func (t *BinaryTree) Inorder() {
	fmt.Println(t.root)
	InorderRecursive(t.root)
}
