package main

import (
	"hello/trie/core"
	"hello/trie/utils"
	"testing"
)

func TestTrie(t *testing.T) {
	t1 := core.NewTrie()
	tt1 := []struct {
		name   string
		toAdd  string
		result bool
	}{
		{"normal trie test 1", "syntax", true},
		{"normal trie test 2", "check", true},
		{"space trie test 1", "syntax check", false},
		{"special character trie test 1", "sample@com", false},
	}
	for _, tc := range tt1 {
		t.Run(tc.name, func(internalT *testing.T) {
			res := t1.Insert(tc.toAdd)
			if res != tc.result {
				internalT.Errorf("failed as output should have been %v but it was %v ", tc.result, res)
			}
		})
	}
	tt2 := []struct {
		name     string
		toSearch string
		result   bool
	}{
		{"normal trie test 1", "syntax", true},
		{"normal trie test 2", "check", true},
		{"space trie test 1", "syntax check", false},
		{"special character trie test 1", "sample@com", false},
	}
	for _, tc := range tt2 {
		t.Run(tc.name, func(internalT *testing.T) {
			res := t1.Search(tc.toSearch)
			if res != tc.result {
				internalT.Errorf("failed as output should have been %v but it was %v ", tc.result, res)
			}
		})
	}
}

func TestOnlyAlphabets(t *testing.T) {
	tt := []struct {
		name   string
		st     string
		result bool
	}{
		{"normal alphabets test 1", "syntax", true},
		{"normal alphabets test 2", "check", true},
		{"space alphabets test 1", "syntax check", false},
		{"special character alphabets test 1", "sample@com", false},
		{"special character alphabets test 1", "sample[com]", false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(internalT *testing.T) {
			res := utils.IsJustAlphabets(tc.st)
			if res != tc.result {
				internalT.Errorf("failed as output should have been %v but it was %v ", tc.result, res)
			}
		})
	}
}

func TestConvertLowerCase(t *testing.T) {
	tt := []struct {
		name string
		st   string
		res  string
	}{
		{"normal lowercase test 1", "syntax", "syntax"},
		{"normal lowercase test 2", "Check", "check"},
		{"space lowercase test 1", "HiGhliGHT", "highlight"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(internalT *testing.T) {
			res := utils.GetCaseInsensitive(tc.st)
			if res != tc.res {
				internalT.Errorf("failed as output should have been %v but it was %v ", tc.res, res)
			}
		})
	}
}
