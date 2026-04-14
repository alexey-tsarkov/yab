package topic

type UUID string

type Lang string

type Topic struct {
	UUID     UUID   `json:"uuid"`
	Title    string `json:"title"`
	Language Lang   `json:"language"`
}

func (t *Topic) String() string {
	return t.Title
}
