package illustrator

type UUID string

type Lang string

type Illustrator struct {
	UUID       UUID   `json:"uuid"`
	Name       string `json:"name"`
	Locale     Lang   `json:"locale"`
	WorksCount int    `json:"works_count"`
}

func (i *Illustrator) String() string {
	return i.Name
}
