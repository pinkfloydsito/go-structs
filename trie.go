package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

func (t *Trie) Insert(word string) {
	var currentNode *TrieNode = t.root

	for idx, c := range word {
		if _, exists := currentNode.children[c]; !exists {
			currentNode.children[c] = NewNode()
		}

		currentNode = currentNode.children[c]

		if idx == len(word)-1 {
			currentNode.isEnd = true
		}
	}
}

func (t *Trie) Search(word string) bool {
	currentNode := t.root

	for _, c := range word {
		if _, exists := currentNode.children[c]; !exists {
			return false
		}

		currentNode = currentNode.children[c]
	}

	return currentNode.isEnd
}

func (t *Trie) PrintTrie() {
	var traverse func(node *TrieNode, prefix string)

	traverse = func(node *TrieNode, prefix string) {
		if node == nil {
			return
		}

		if node.isEnd {
			fmt.Println("Word:", prefix)
		}

		for char, childNode := range node.children {
			traverse(childNode, prefix+string(char))
		}
	}

	traverse(t.root, "")
}
