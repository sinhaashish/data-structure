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

	fmt.Println("PreOrder Traversal - recursive solution : ")
	binaryTree.PreOrder()

	fmt.Println("PreOrder Traversal - Iterative solution : ")
	binaryTree.PreOrderIterative()
	fmt.Println()

	fmt.Println("Inorder Traversal - recursive solution : ")
	binaryTree.InOrder()

	fmt.Println("Inorder Traversal - recursive solution : ")
	binaryTree.InOrderIterative()

	fmt.Println("PostOrder Traversal - recursive solution : ")
	binaryTree.PostOrder()

	fmt.Println("PostOrder Traversal - recursive solution : ")
	//	binaryTree.PostOrderIterative()

	fmt.Println("Level Order trasversal: ")
	binaryTree.LevelOrder()
	fmt.Println()

}
