package Day11

import "testing"

func TestDay11(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"default"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Day11()
		})
	}
}