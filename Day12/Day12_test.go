package Day12

import "testing"

func TestDay12(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"default"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Day12()
		})
	}
}
