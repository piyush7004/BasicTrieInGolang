package dict

//TrieNode A generic node for our trie
type TrieNode struct {
	isEnd    bool
	children map[byte]TrieNode
}

var root TrieNode

//NewTrie Use this to retrieve the root trienode
func NewTrie() *TrieNode {
	return &root
}

//Insert will add a string 's' to this trie
func (node TrieNode) Insert(s string) {

}

//Search will search a string to this trie and return true/false accordingly
func (node TrieNode) Search(s string) bool {
	return false
}
