package searchengine

import (
	"strings"
)

type InvertedIndex struct {
	index map[string]map[Document]struct{} // map of word and set of documents
	trie *Trie
}


// new index
func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		index: make(map[string]map[Document]struct{}),
		trie: NewTrie(),
	}
}

// add document in index
func (i *InvertedIndex) AddDocument(document Document) {
	words := strings.Split(strings.ToLower(document.content), "\\W+")
	for _, word := range words {
		if _, exists := i.index[word]; !exists {
			set := make(map[Document]struct{})
			set[document] = struct{}{}
			i.index[word] = set
		} else {
			i.index[word][document] = struct{}{}
		}
		i.trie.Insert(word)
	}
}
// get documents
func(i *InvertedIndex) GetDocumentsForWord(word string) map[Document]struct{} {
	return i.index[word]
}

