package main

import (
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	trie := &Trie{
		root: NewNode(),
	}

	trie.Insert("Hello")

	tests := []struct {
		word     string
		expected bool
	}{
		{"Hello", true},
		{"hello", false},
	}

	for _, test := range tests {
		result := trie.Search(test.word)

		if result != test.expected {
			t.Errorf("expected %t got %t", test.expected, result)
		}
	}
}
