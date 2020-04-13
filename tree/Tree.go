package main

import (
	"fmt"
	"tree/util"
)

func main() {

	binaryTree := &util.BinaryTree{}
	binaryTree.Insert(4).
		Insert(2).
		Insert(6).
		Insert(1).
		Insert(3).
		Insert(5).
		Insert(7)
	fmt.Println("Inorder Traversal - recursive solution : ")

	binaryTree.Inorder()
	fmt.Println()

}
