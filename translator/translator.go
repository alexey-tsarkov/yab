package translator

type UUID string

type Lang string

type Translator struct {
	UUID       UUID   `json:"uuid"`
	Name       string `json:"name"`
	Locale     Lang   `json:"locale"`
	WorksCount int    `json:"works_count"`
}

func (t *Translator) String() string {
	return t.Name
}
