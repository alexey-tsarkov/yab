package topic

type Topic struct {
	UUID     string `json:"uuid"`
	Title    string `json:"title"`
	Language string `json:"language"`
}

func (t *Topic) String() string {
	return t.Title
}
