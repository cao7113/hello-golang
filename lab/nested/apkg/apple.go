package apkg

type Apple struct {
	Kind  string
	Color string
}

func (a *Apple) SetColor(col string) {
	a.Color = col
}
