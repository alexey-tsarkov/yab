package internal

type Illustrator struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Locale     string `json:"locale"`
	WorksCount int    `json:"works_count"`
}

func (i *Illustrator) String() string {
	return i.Name
}
