package util

import (
	"fmt"
)

// PreOrderRecursive : Inorder recursive traversal
func PreOrderRecursive(root *Node) {
	if root != nil {
		fmt.Printf("%d ->  ", root.data)
		PreOrderRecursive(root.left)
		PreOrderRecursive(root.right)
	}
}

// Inorder : traverses the tree
func (t *BinaryTree) Inorder() {
	PreOrderRecursive(t.root)
}
