package searchengine

import "time"

type Document struct {
	name      string
	content   string
	author    string
	createdAt time.Time
	updatedAt time.Time
}

func NewDocument(name, content, author string) *Document {
	return &Document{
		name: name,
		content: content,
		author: author,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}