package enumer

//first run: go get -u golang.org/x/tools/...
//go:generate stringer -type=Gender

type Gender int

const (
	Male Gender = iota
	Female
)
