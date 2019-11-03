package core

import (
	"fmt"
	"os"
	"strconv"
)

//BinarySearchTreeNode represents a typical binary tree node
type BinarySearchTreeNode struct {
	data       int
	leftNode   *BinarySearchTreeNode
	rightNode  *BinarySearchTreeNode
	parentNode *BinarySearchTreeNode
}

var (
	root       BinarySearchTreeNode
	leftChild  *BinarySearchTreeNode
	rightChild *BinarySearchTreeNode
)

const (
	//InOrder signifies the inorder traversal
	InOrder = "inorder"
	//PreOrder signifies the preorder traversal
	PreOrder = "preorder"
	//PostOrder signifies the postorder traversal
	PostOrder = "postorder"
)

//NewBinarySearchTree returns the address to root of binary tree
func NewBinarySearchTree(rootData int) *BinarySearchTreeNode {
	root = BinarySearchTreeNode{
		data:       rootData,
		leftNode:   nil,
		rightNode:  nil,
		parentNode: nil,
	}
	return &root
}

//Insert adds an element to the tree
func (root *BinarySearchTreeNode) Insert(data int) *BinarySearchTreeNode {
	fmt.Println("inserting ", data)
	current := insertNode(root, data)
	return current
}

//Find finds an entry in the bst and returns its iterations else -1
func (root *BinarySearchTreeNode) Find(data int) int {
	var iter int
	iter = search(root, data, iter)
	return iter
}

func search(node *BinarySearchTreeNode, data int, iter int) int {
	if node == nil {
		return -1
	}
	if data == node.data {
		return iter
	} else if data < node.data {
		iter = search(node.leftNode, data, iter+1)
	} else {
		iter = search(node.rightNode, data, iter+1)
	}
	return iter
}

//Save sasves the current tree in the requested format to a local file
func (root *BinarySearchTreeNode) Save(order string) {
	f, err := os.Create("tree")
	if err != nil {
		fmt.Println(err)
	}
	switch order {
	case InOrder:
		traverseInOrder(root, f)
	case PreOrder:
		traversePreOrder(root, f)
	case PostOrder:
		traversePostOrder(root, f)
	}
	// fmt.Println()
}

func insertNode(node *BinarySearchTreeNode, data int) *BinarySearchTreeNode {
	if node == nil {
		return &BinarySearchTreeNode{
			data:       data,
			leftNode:   nil,
			rightNode:  nil,
			parentNode: nil,
		}
	}
	if data < node.data {
		leftChild = insertNode(node.leftNode, data)
		node.leftNode = leftChild
		leftChild.parentNode = node
	} else {
		rightChild = insertNode(node.rightNode, data)
		node.rightNode = rightChild
		rightChild.parentNode = node
	}

	return node

}

func traverseInOrder(node *BinarySearchTreeNode, f *os.File) {
	if node == nil {
		return
	}
	traverseInOrder(node.leftNode, f)
	f.WriteString(strconv.Itoa(node.data))
	traverseInOrder(node.rightNode, f)
}

func traversePreOrder(node *BinarySearchTreeNode, f *os.File) {
	if node == nil {
		return
	}
	f.WriteString(strconv.Itoa(node.data))
	// fmt.Printf("%d ", node.data)
	traversePreOrder(node.leftNode, f)
	traversePreOrder(node.rightNode, f)
}

func traversePostOrder(node *BinarySearchTreeNode, f *os.File) {
	if node == nil {
		return
	}
	traversePostOrder(node.leftNode, f)
	traversePostOrder(node.rightNode, f)
	f.WriteString(strconv.Itoa(node.data))
	// fmt.Printf("%d ", node.data)
}
