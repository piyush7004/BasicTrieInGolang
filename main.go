package main

import (
	"fmt"
	"hello/trie/core"
)

var t *core.TrieNode

func init() {
	t = core.NewTrie()
}

func main() {
	t.Insert("abcd")
	t.Insert("lmop")
	t.Insert("lmop")
	fmt.Println("abcd-", t.Search("abcd"))
	fmt.Println("abc-", t.Search("abc"))
	fmt.Println("abcd-", t.Search("abcd"))
	t.Insert("lmno")
	fmt.Println("lmnop-", t.Search("lmnop"))
	fmt.Println("lm-", t.Search("lm"))
	fmt.Println("lmno-", t.Search("lmno"))
	t.Insert("23asd")
	fmt.Println("23asd-", t.Search("23asd"))
	fmt.Println("Abcd-", t.Search("Abcd"))
	t.Insert("syntax check")
	fmt.Println("syntax check", t.Search("syntax check"))

}
