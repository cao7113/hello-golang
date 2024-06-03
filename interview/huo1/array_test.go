package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type node struct {
	key int
	val int
}

func (s *ASuite) TestLoopVar() {
	list := []node{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 5},
		{7, 5},
	}
	result := make([]*node, 0, 1)
	for _, nd := range list {
		if nd.val > 1 {
			result = append(result, &nd)
		}
	}
	var sum int
	for _, el := range result {
		sum += el.val
	}

	s.Equal(25, sum)
	s.Equal(8, cap(result))
}

func TestASuite(t *testing.T) {
	suite.Run(t, &ASuite{})
}

type ASuite struct {
	suite.Suite
}
