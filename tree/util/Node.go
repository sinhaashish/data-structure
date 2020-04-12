package util

// Node for tree element
type Node struct {
	data  int
	left  *Node
	right *Node
}

// BinaryTree creates a Binary tree
type BinaryTree struct {
	root *Node
}

// Insert : insert the node in Binary tree
func (t *BinaryTree) Insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &Node{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

// Insert : Insert the node in the Binary tree
func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}
