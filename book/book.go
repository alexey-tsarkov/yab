package book

import "github.com/alexey-tsarkov/yab/audio"

type UUID string

type Language string

const (
	EN = Language("en")
	RU = Language("ru")
)

type Book struct {
	UUID             UUID           `json:"uuid"`
	Title            string         `json:"title"`
	Language         Language       `json:"language"`
	Year             int            `json:"original_year"`
	Annotation       string         `json:"annotation"`
	Authors          []*Author      `json:"authors_objects"`
	Translators      []*Translator  `json:"translators"`
	Publishers       []*Publisher   `json:"publishers"`
	Illustrators     []*Illustrator `json:"illustrators"`
	Pages            int            `json:"paper_pages"`
	Topics           []*Topic       `json:"topics"`
	DocumentUUID     UUID           `json:"document_uuid"`
	LinkedAudioBooks []audio.UUID   `json:"linked_audiobook_uuids"`
}

type Topic struct {
	UUID     UUID     `json:"uuid"`
	Title    string   `json:"title"`
	Language Language `json:"language"`
}

type Author = entity

type Translator = entity

type Publisher = entity

type Illustrator = entity

type entity struct {
	UUID       UUID     `json:"uuid"`
	Name       string   `json:"name"`
	Locale     Language `json:"locale"`
	WorksCount int      `json:"works_count"`
}

func (b *Book) String() string {
	return b.Title
}

func (t *Topic) String() string {
	return t.Title
}

func (e *entity) String() string {
	return e.Name
}
