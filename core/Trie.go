package core

import (
	"fmt"
	"hello/trie/utils"
	"reflect"
)

//TrieNode A generic node for our trie
type TrieNode struct {
	isEnd    bool
	contents map[byte]*TrieNode
}

var root TrieNode

// var m map[byte]TrieNode

//NewTrie Use this to retrieve the root trienode
func NewTrie() *TrieNode {
	m := make(map[byte]*TrieNode)
	root = TrieNode{
		isEnd:    false,
		contents: m,
	}
	return &root
}

//Insert will add a string 's' to this trie
func (node *TrieNode) Insert(s string) bool {
	fmt.Println("inserting ", s)
	if !utils.IsJustAlphabets(s) {
		fmt.Println("wrong input provided")
		return false
	}
	s = utils.GetCaseInsensitive(s)
	c := 0
	current := node
	// var tmp1 *TrieNode
	for _, b := range []byte(s) {
		_, found := current.contents[b]
		if found {
			current = current.contents[b]
			// tmp1 = &current
		} else {
			tmp := TrieNode{false, make(map[byte]*TrieNode)}
			current.contents[b] = &tmp
			current = &tmp
			// tmp1 = &current
		}
		c++
	}
	current.isEnd = true
	return true
}

//Search will search a string to this trie and return true/false accordingly
func (node *TrieNode) Search(s string) bool {
	fmt.Println("searching ", s)
	// fmt.Println(node)
	if !utils.IsJustAlphabets(s) {
		fmt.Println("wrong input for search")
		return false
	}
	s = utils.GetCaseInsensitive(s)
	var tmp *TrieNode
	current := node
	c := 0
	for _, b := range []byte(s) {
		// fmt.Println(c, b)
		_, found := current.contents[b]
		if !found {
			fmt.Println("this word isn't present yet")
			return false
		}

		tmp = current.contents[b]
		if reflect.DeepEqual(tmp, (TrieNode{})) {
			fmt.Println("returning from here")
			return false
		}
		current = tmp
		c++
	}
	return current.isEnd
}
