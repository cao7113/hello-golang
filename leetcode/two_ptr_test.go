package leetcode

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/sum-of-square-numbers/
func (s *LtCodeSuite) TestP633() {
	cases := []struct {
		num  int
		want bool
	}{
		{5, true},
		{4, true},
		{3, false},
		{2, true},
		{1, true},
	}
	for _, c := range cases {
		s.Run(fmt.Sprintf("%+v", c), func() {
			s.Equal(c.want, judgeSquareSum(c.num))
		})
	}
}

func judgeSquareSum(c int) bool {
	s := math.Sqrt(float64(c))
	l, r := 0, int(s)
	for l <= r {
		sum := l*l + r*r
		if sum == c {
			return true
		}
		if sum < c {
			l++
		} else {
			l--
		}
	}
	return false
}
