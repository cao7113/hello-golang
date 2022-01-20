package leetcode

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestLtCodeSuite(t *testing.T) {
	suite.Run(t, &LtCodeSuite{})
}

type LtCodeSuite struct {
	suite.Suite
}
