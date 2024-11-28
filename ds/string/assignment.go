package string

import "sort"

type SortedString []rune

func (s SortedString) Len() int {
	return len(s)
}

func (s SortedString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortedString) Less(i, j int) bool {
	return s[i] < s[j]
}

func ClubAnagramWord(words []string) [][]string {
	wordMap := make(map[string][]string)
	var ans [][]string
	for _, word := range words {
		charWordSortInterface := SortedString([]rune(word))
		sort.Sort(charWordSortInterface)
		sortedWord := string(charWordSortInterface)
		if values,  exists := wordMap[sortedWord]; exists {
			wordMap[sortedWord] = append(values, word)
		} else {
			wordMap[sortedWord] = append([]string{}, word)
		}
	}
	for _, values := range wordMap {
		ans = append(ans, values)
	}
	return ans
}