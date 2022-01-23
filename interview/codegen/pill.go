package painkiller

// run: go install golang.org/x/tools/...@latest

//go:generate stringer -type=Pill

type Pill int

const (
	A Pill = iota // 0
	B             // 1
	C             // 2
	D = C         // 2
	E             // 2
	F             // 2
)
