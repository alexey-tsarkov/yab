package internal

type Narrator struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Locale     string `json:"locale"`
	WorksCount int    `json:"works_count"`
}

func (n *Narrator) String() string {
	return n.Name
}
