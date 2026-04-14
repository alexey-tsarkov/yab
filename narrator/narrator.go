package narrator

type UUID string

type Lang string

type Narrator struct {
	UUID       UUID   `json:"uuid"`
	Name       string `json:"name"`
	Locale     Lang   `json:"locale"`
	WorksCount int    `json:"works_count"`
}

func (n *Narrator) String() string {
	return n.Name
}
