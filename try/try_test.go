package try

import "testing"

func TestTry(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test1"},
		{name: "test2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Try()
		})
	}
}
