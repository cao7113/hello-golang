package interview

import (
	"github.com/stretchr/testify/suite"
	"math/rand"
	"sort"
	"testing"
)

func (s *SortSuite) TestStableSort() {
	a := []int{3, 1, 2}
	sort.Ints(a)
	s.Equal([]int{1, 2, 3}, a)
}

// todo more sort try

func (s *SortSuite) TestQuickSort() {
	cases := []struct {
		input []int
		want  []int
	}{
		{[]int{3, 2, 1, 4, 8}, []int{1, 2, 3, 4, 8}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
	}
	for _, it := range cases {
		s.Equal(it.want, quickSort(it.input))
	}
}

func quickSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}
	base := rand.Intn(n)
	el := a[base]

	var left, right, same []int
	for _, it := range a {
		if it == el {
			same = append(same, it)
		} else if it > el {
			right = append(right, it)
		} else {
			left = append(left, it)
		}
	}
	left = quickSort(left)
	right = quickSort(right)

	left = append(left, same...)
	return append(left, right...)
}

func TestSortSuite(t *testing.T) {
	suite.Run(t, &SortSuite{})
}

type SortSuite struct {
	suite.Suite
}
