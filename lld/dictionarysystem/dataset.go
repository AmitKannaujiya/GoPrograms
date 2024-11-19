package dictionarysystem

type Dataset struct {
	datasetName string
	trie *Trie
}

func NewDataset(datasetName string) *Dataset {
	return &Dataset{
		datasetName: datasetName,
		trie: NewTrie(),
	}
}

func (d *Dataset) IsPresent(word string) bool {
	return d.trie.Search(word)
}

func (d *Dataset) AddWord(word string) {
	d.trie.Insert(word)
}