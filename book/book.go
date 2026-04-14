package book

import (
	"github.com/alexey-tsarkov/yab/audio"
	"github.com/alexey-tsarkov/yab/author"
	"github.com/alexey-tsarkov/yab/illustrator"
	"github.com/alexey-tsarkov/yab/publisher"
	"github.com/alexey-tsarkov/yab/topic"
	"github.com/alexey-tsarkov/yab/translator"
)

type UUID string

type Lang string

type Book struct {
	UUID             UUID                       `json:"uuid"`
	Title            string                     `json:"title"`
	Language         Lang                       `json:"language"`
	Year             int                        `json:"original_year"`
	Annotation       string                     `json:"annotation"`
	Authors          []*author.Author           `json:"authors_objects"`
	Translators      []*translator.Translator   `json:"translators"`
	Publishers       []*publisher.Publisher     `json:"publishers"`
	Illustrators     []*illustrator.Illustrator `json:"illustrators"`
	Pages            int                        `json:"paper_pages"`
	Topics           []*topic.Topic             `json:"topics"`
	LinkedAudioBooks []audio.UUID               `json:"linked_audiobook_uuids"`
}

func (b *Book) String() string {
	return b.Title
}
