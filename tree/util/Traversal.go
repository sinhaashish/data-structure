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

// PreOrderIterative : ProOrder in Iterative fashion
func (t *BinaryTree) PreOrderIterative() {
	var st Stack
	st.Push(*t.root)
	for {
		if st.IsEmpty() {
			break
		} else {
			currentNode, _ := st.Pop()
			fmt.Printf("%d -> ", currentNode.data)
			if currentNode.right != nil {
				st.Push(*currentNode.right)
			}
			if currentNode.left != nil {
				st.Push(*currentNode.left)
			}
		}
	}
}

// InOrderIterative - Traverses the three in inorder fashion
func (t *BinaryTree) InOrderIterative() {
	var st Stack
	var done bool
	currentNode := t.root
	for {
		if done {
			break
		}
		if currentNode != nil {
			st.Push(*currentNode)
			currentNode = currentNode.left
		} else {
			if st.IsEmpty() {
				done = true
			} else {
				*currentNode, _ = st.Pop()
				fmt.Printf("%d -> ", currentNode.data)
				currentNode = currentNode.right
			}
		}
	}
}

// PostOrderIterative - Traverses the three in postorder fashion
// func (t *BinaryTree) PostOrderIterative() {
// 	var st Stack
// 	st.Push(*t.root)
// 	for {
// 		if st.IsEmpty() {
// 			break
// 		} else {
// 			node, _ := st.Peek()
// 			// fmt.Printf("%d -> ", node.data)

// 			if node.right != nil {
// 				st.Push(*node.right)
// 			}
// 			if node.left != nil {
// 				st.Push(*node.left)
// 			}
// 		}
// 	}
// }

// LevelOrder : Traverses the tree in level order
func (t *BinaryTree) LevelOrder() {
	var q Queue
	q.Push(*t.root)
	for {
		if q.IsEmpty() {
			break
		} else {
			node, _ := q.Pop()
			fmt.Printf("%d -> ", node.data)

			if node.left != nil {
				q.Push(*node.left)
			}
			if node.right != nil {
				q.Push(*node.right)
			}
		}
	}
}
