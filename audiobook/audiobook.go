package audiobook

import (
	"github.com/alexey-tsarkov/yab/author"
	"github.com/alexey-tsarkov/yab/narrator"
	"github.com/alexey-tsarkov/yab/publisher"
	"github.com/alexey-tsarkov/yab/topic"
	"github.com/alexey-tsarkov/yab/translator"
)

type AudioBook struct {
	UUID        string                   `json:"uuid"`
	Title       string                   `json:"title"`
	Language    string                   `json:"language"`
	Year        int                      `json:"original_year"`
	Annotation  string                   `json:"annotation"`
	Authors     []*author.Author         `json:"authors_objects"`
	Translators []*translator.Translator `json:"translators"`
	Publishers  []*publisher.Publisher   `json:"publishers"`
	Narrators   []*narrator.Narrator     `json:"narrators"`
	Duration    int                      `json:"duration"`
	Topics      []*topic.Topic           `json:"topics"`
	LinkedBooks []string                 `json:"linked_book_uuids"`
}

func (ab *AudioBook) String() string {
	return ab.Title
}
