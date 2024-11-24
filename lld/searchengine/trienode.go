package searchengine

type TrieNode struct {
	char     rune
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			char: '-',
			children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		if _, exists :=  current.children[char]; !exists {
			current.children[char] = &TrieNode{
				char: char,
				children: make(map[rune]*TrieNode),
			}
		}
		current = current.children[char]
	}
	current.isWord = true
}

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