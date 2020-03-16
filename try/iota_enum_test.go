package try

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

// Iota usage
const (
	Male   = iota // 0
	Female        // 1
)

const (
	_  = iota + 1 // 1
	_             // skip
	C3            // 3
)

func TestIota(t *testing.T) {
	assert.Equal(t, Male, 0)
	assert.Equal(t, Female, 1)

	assert.Equal(t, C3, 3)
	assert.Equal(t, C3, 3)
}

// Iota used in enum like type
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func TestEnum(t *testing.T) {
	d := East
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
	assert.Equal(t, West.String(), "West")
}

// Another common application for iota is to implement a bitmask, a small set of booleans, often called flags,
// represented by the bits in a single number.
// https://yourbasic.org/golang/bitmask-flag-set-clear/
// ref log.LstdFlags

/*

https://yourbasic.org/golang/iota/
The iota keyword represents successive integer constants 0, 1, 2,…
It resets to 0 whenever the word const appears in the source code,
and increments after each const specification.

Here we rely on the fact that expressions are implicitly repeated in a paren­thesized const declaration
– this indicates a repetition of the preceding expression and its type.
same as
const (
	Male   = iota // 0
	Female = iota // 1
)
*/
