package main

import (
	tree "github.com/SeiyaJapon/max_depth_binary_tree/tree"
)

func main() {
	myTree := tree.NewBinaryTree()

	myTree.Insert(3, false, nil)
	myTree.Insert(9, true, myTree.Root)
	myTree.Insert(4, false, myTree.Search(9))
	myTree.Insert(20, false, myTree.Root)
	myTree.Insert(15, true, myTree.Search(20))
	myTree.Insert(7, false, myTree.Search(20))
	myTree.Insert(1, false, myTree.Search(7))

	myTree.PrintTree()

	maxDepth := myTree.MaxDepth()
	println("Max Depth of the tree:", maxDepth)
}
