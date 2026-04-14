package translator

type Translator struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Locale     string `json:"locale"`
	WorksCount int    `json:"works_count"`
}

func (t *Translator) String() string {
	return t.Name
}
