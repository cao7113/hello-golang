package trytest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	fmt.Println("Standard library testing style with table-driven")
	type args struct {
		num int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"test1", args{1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := AddOne(tt.args.num); gotResult != tt.wantResult {
				t.Errorf("AddOne() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestAddOneWithTestify(t *testing.T) {
	assert.Equal(t, 2, AddOne(1))
}

func TestAddOneWithTestifyBetter(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, AddOne(1))
}
