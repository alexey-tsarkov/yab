package book

import (
	. "github.com/alexey-tsarkov/yab/internal"
)

type Book struct {
	UUID             string         `json:"uuid"`
	Title            string         `json:"title"`
	Language         string         `json:"language"`
	Year             int            `json:"original_year"`
	Annotation       string         `json:"annotation"`
	Authors          []*Author      `json:"authors_objects"`
	Translators      []*Translator  `json:"translators"`
	Publishers       []*Publisher   `json:"publishers"`
	Illustrators     []*Illustrator `json:"illustrators"`
	Pages            int            `json:"paper_pages"`
	Topics           []*Topic       `json:"topics"`
	LinkedAudioBooks []string       `json:"linked_audiobook_uuids"`
}

func (b *Book) String() string {
	return b.Title
}
