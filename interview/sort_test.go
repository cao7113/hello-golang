package interview

import (
	"github.com/stretchr/testify/suite"
	"sort"
	"testing"
)

func (s *SortSuite) TestStableSort() {
	a := []int{3, 1, 2}
	sort.Ints(a)
	s.Equal([]int{1, 2, 3}, a)
}

func (s *SortSuite) TestQuickSort() {
	// todo
}

func quickSort(a []int) []int {
	var r []int
	//for _, e := range a {
	// todo
	//}
	return r
}

func TestSortSuite(t *testing.T) {
	suite.Run(t, &SortSuite{})
}

type SortSuite struct {
	suite.Suite
}
