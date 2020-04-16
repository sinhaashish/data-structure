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

// PreOrder : traverses the tree in Preprder
func (t *BinaryTree) PreOrder() {
	PreOrderRecursive(t.root)
	fmt.Println()
}

// InOrderRecursive : Inorder recursive traversal
func InOrderRecursive(root *Node) {
	if root != nil {
		InOrderRecursive(root.left)
		fmt.Printf("%d ->  ", root.data)
		InOrderRecursive(root.right)
	}
}

// InOrder : traverses the tree in InOrder
func (t *BinaryTree) InOrder() {
	InOrderRecursive(t.root)
	fmt.Println()
}

// PostrderRecursive : Post order recursive traversal
func PostrderRecursive(root *Node) {
	if root != nil {
		PostrderRecursive(root.left)
		PostrderRecursive(root.right)
		fmt.Printf("%d ->  ", root.data)
	}
}

// PostOrder : traverses the tree in PostOrder
func (t *BinaryTree) PostOrder() {
	PostrderRecursive(t.root)
	fmt.Println()
}
