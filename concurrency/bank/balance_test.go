package bank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_teller(t *testing.T) {
	Deposit(20)
	b := Balance()
	assert.Equal(t, 20, b)
}
