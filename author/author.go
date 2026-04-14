package author

type UUID string

type Lang string

type Author struct {
	UUID       UUID   `json:"uuid"`
	Name       string `json:"name"`
	Locale     Lang   `json:"locale"`
	WorksCount int    `json:"works_count"`
}

func (a *Author) String() string {
	return a.Name
}
