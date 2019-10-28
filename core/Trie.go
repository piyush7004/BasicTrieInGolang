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

//Delete will delete a provided word and also restructure the trie, if required
func (node *TrieNode) Delete(s string) bool {
	fmt.Println("Deleting ", s)
	//check if provided string is all alphabets
	if !utils.IsJustAlphabets(s) {
		return false
	}
	//get the all lower case equivalent of that, if not already
	s = utils.GetCaseInsensitive(s)
	current := node
	return deleteInternal(current, s, 0)
}

//deleteInternal is a recursive function as it is going to perform
//deletion of nodes backwards too based on the size of the deleted node
func deleteInternal(current *TrieNode, s string, i int) bool {
	//signifies the traversal to the last element of the trie, and the recursion unfolds from there
	if i == len(s) {
		/*
			if isEnd of current is false at this point (when we are at the end of the word),
			it means that word 's' is present in this trie but as a prefix and not as a word. In other words,
			the "word" s is not present in the trie  therefore it is good idea to just
			keep that as is and return from that point.
		*/
		if !current.isEnd {
			return false
		}
		/*
			isEnd true signifies that "word" s is present and it is right candidate to be deleted
			however we need to make sure that isEnd it getting off (false). Also return true if
			it has no other characters in it to signify that this node is a candidate of deletion now
		*/
		current.isEnd = false
		if len(current.contents) == 0 {
			return true
		} else {
			return false
		}
	}
	//if at any certain point, the node pointing to the current-character of the word is not existing
	b := []byte(s)[i]
	tmp, found := current.contents[b]
	if !found {
		return false
	}
	// fmt.Println(tmp, i)
	//if we are here, it means tmp is a valid node, so let's loop back using it and the next index
	canCurrentNodeBeDeleted := deleteInternal(tmp, s, i+1)
	if canCurrentNodeBeDeleted {
		delete(current.contents, b)
		if len(current.contents) == 0 && !current.isEnd {
			return true
		} else {
			return false
		}
	}

	return false
}
