package internal

type Publisher struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Locale     string `json:"locale"`
	WorksCount int    `json:"works_count"`
}

func (p *Publisher) String() string {
	return p.Name
}
