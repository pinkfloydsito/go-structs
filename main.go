package main

import "fmt"

func main() {
	trie := &Trie{
		root: NewNode(),
	}

	trie.Insert("Hola")
	trie.Insert("Mundo")

	trie.PrintTrie()
	fmt.Println(trie.Search("Hola"))
}
