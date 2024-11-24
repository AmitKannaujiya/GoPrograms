package searchengine

type Category struct {
	categoryName string
	documents []Document
	index *InvertedIndex
}


func NewCategory(categoryName string) *Category {
	return &Category{
		categoryName: categoryName,
		documents: make([]Document, 0),
		index: NewInvertedIndex(),
	}
}

func(c *Category) AddDocument(document Document) {
	c.documents = append(c.documents, document)
	c.UpdateInvertedIndex(document)
}

func(c *Category) UpdateInvertedIndex(document Document) {
	c.index.AddDocument(document)
}