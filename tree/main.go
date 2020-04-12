package main

import (
	"fmt"
	"tree/util"
)

func main() {

	binaryTree := &util.BinaryTree{}
	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(4)
	binaryTree.Insert(5)
	binaryTree.Insert(6)
	binaryTree.Insert(7)
	binaryTree.Insert(8)

	fmt.Println("Inorder Traversal - recursive solution : ")

	InorderRecursive(binaryTree)

}
