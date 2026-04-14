package publisher

type UUID string

type Lang string

type Publisher struct {
	UUID       UUID   `json:"uuid"`
	Name       string `json:"name"`
	Locale     Lang   `json:"locale"`
	WorksCount int    `json:"works_count"`
}

func (p *Publisher) String() string {
	return p.Name
}
