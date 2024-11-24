package searchengine

import (
	"strings"
)

type ISearchStretegy interface {
	Search(Category, string) []Document 
}

type UnorderSearch struct {}
type OrderSearch struct {}
type SubstringSearch struct {}
// Unorder search
func(us *UnorderSearch) Search(category Category, pattern string) []Document {
	words := strings.Split(pattern, " ")
	documents := category.index.GetDocumentsForWord(words[0])
	var result [] Document

	for i:= 1; i < len(words); i++ {
		docs := category.index.GetDocumentsForWord(words[i])
		for document, _ :=  range docs {
			if _, exists := documents[document]; !exists {
				documents[document] = struct{}{}
			}
		}
	}
	for document, _ := range documents {
		result = append(result, document)
	}
	return result
}
// substrings search
func(us *SubstringSearch) Search(category Category, pattern string) []Document {
	word := strings.ToLower(pattern)
	var result [] Document
	for _, document := range category.documents {
		if strings.Contains(document.content, word) {
			result = append(result, document)
		}
	}
	return result
}
// order search
func(os *OrderSearch) Search(category Category, pattern string) []Document {
	word := strings.ToLower(pattern)
	var result [] Document
	for _, document := range category.documents {
		if category.index.trie.Search(word){
			result = append(result, document)
		}
	}
	return result
}
