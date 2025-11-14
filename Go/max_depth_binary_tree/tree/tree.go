package tree

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

func (bt *BinaryTree) Insert(val int, isLeft bool, node *TreeNode) {
	if bt.Root == nil {
		bt.Root = &TreeNode{Val: val}
		return
	}

	if node == nil {
		node = bt.Root
	}

	insertNode(node, val, isLeft)
}

func insertNode(current *TreeNode, val int, isLeft bool) {
	node := findNode(current, val)
	if node == nil {
		node = current
	}

	if isLeft {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
		}
	}
}

func (bt *BinaryTree) Search(val int) *TreeNode {
	return findNode(bt.Root, val)
}

func findNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val == val {
		return node
	}

	LeftResult := findNode(node.Left, val)
	if LeftResult != nil {
		return LeftResult
	}

	return findNode(node.Right, val)
}

func (bt *BinaryTree) MaxDepth() int {
	return findMaxDepth(bt.Root, 0)
}

func findMaxDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	depth++

	leftDepth := findMaxDepth(node.Left, depth)
	rightDepth := findMaxDepth(node.Right, depth)

	if leftDepth > rightDepth {
		return leftDepth
	}

	return rightDepth
}

func (bt *BinaryTree) GetNodeByValue(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val == val {
		return node
	}

	result := &TreeNode{}

	if node.Left != nil {
		result = bt.GetNodeByValue(node.Left, val)
	}
	if node.Right != nil {
		if result == nil {
			result = bt.GetNodeByValue(node.Right, val)
		}
	}

	return result
}

func CountNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + CountNodes(node.Left) + CountNodes(node.Right)
}

func (bt *BinaryTree) PrintTree() {
	printTree(bt.Root, 0)
}

func printTree(node *TreeNode, level int) {
	if node == nil {
		return
	}

	indent := strings.Repeat("    ", level)
	fmt.Printf("%s%d\n", indent, node.Val)

	printTree(node.Left, level+1)
	printTree(node.Right, level+1)
}
