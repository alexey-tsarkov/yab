package audiobook

import (
	. "github.com/alexey-tsarkov/yab/internal"
)

type AudioBook struct {
	UUID        string        `json:"uuid"`
	Title       string        `json:"title"`
	Language    string        `json:"language"`
	Year        int           `json:"original_year"`
	Annotation  string        `json:"annotation"`
	Authors     []*Author     `json:"authors_objects"`
	Translators []*Translator `json:"translators"`
	Publishers  []*Publisher  `json:"publishers"`
	Narrators   []*Narrator   `json:"narrators"`
	Duration    int           `json:"duration"`
	Topics      []*Topic      `json:"topics"`
	LinkedBooks []string      `json:"linked_book_uuids"`
}

func (ab *AudioBook) String() string {
	return ab.Title
}
