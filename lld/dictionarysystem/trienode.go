package dictionarysystem

type TrieNode struct {
	children map[rune]*TrieNode
	char     rune
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			char:     '-', // root char
			isWord:   false,
		},
	}
}

// insert into trie
func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				char:     char,
			}
		}
		current = current.children[char]
	}
	current.isWord = true
}

// search in trie
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}
	return current.isWord
}