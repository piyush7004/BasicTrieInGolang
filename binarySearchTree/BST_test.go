package main

import (
	"hello/DSAinGoLang/binarySearchTree/core"
	"io/ioutil"
	"testing"
)

func TestInsertAndShow(t *testing.T) {
	var root *core.BinarySearchTreeNode
	var bytes []byte
	tt := []struct {
		name     string
		toInsert int
		tree     string
	}{
		{"test insert 1", 3, "3"},
		{"test insert 2", 1, "13"},
		{"test insert 3", 2, "123"},
		{"test insert 4", 6, "1236"},
	}
	for i, tc := range tt {
		t.Run(tc.name, func(internalT *testing.T) {
			if i == 0 {
				root = core.NewBinarySearchTree(tc.toInsert)
			} else {
				root.Insert(tc.toInsert)
			}
			root.Save(core.InOrder)
			bytes, _ = ioutil.ReadFile("tree")
			if string(bytes) != tc.tree {
				internalT.Errorf("failed as received value should have been %s but it is %s", tc.tree, string(bytes))
			}
		})
	}
	// os.Remove("tree")
}

func TestInsertAndFind(t *testing.T) {
	var root *core.BinarySearchTreeNode
	tt := []struct {
		name       string
		toInsert   int
		iterations int
	}{
		{"test find 1", 3, 0},
		{"test find 2", 2, 1},
		{"test find 3", 1, 2},
		{"test find 4", 6, 1},
	}
	for i, tc := range tt {
		t.Run(tc.name, func(internalT *testing.T) {
			if i == 0 {
				root = core.NewBinarySearchTree(tc.toInsert)
			} else {
				root.Insert(tc.toInsert)
			}
			// root.Save(core.InOrder)
			if root.Find(tc.toInsert) != tc.iterations {
				internalT.Errorf("failed as received value should have been %d but it is %d", tc.iterations, root.Find(tc.toInsert))
			}
		})
	}
}
