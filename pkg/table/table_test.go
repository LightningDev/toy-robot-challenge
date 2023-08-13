package table

import (
	"strconv"
	"testing"
)

func TestIsValidPosition(t *testing.T) {
	tests := []struct {
		x, y     int
		expected bool
	}{
		// Inside the board
		{2, 2, true},
		{0, 0, true},
		{4, 4, true},

		// On the boundaries
		{5, 2, false},
		{2, 5, false},
		{5, 5, false},
		{-1, 2, false},
		{2, -1, false},

		// Outside the board
		{6, 2, false},
		{2, 6, false},
		{-2, 2, false},
		{2, -2, false},
	}

	for _, tt := range tests {
		t.Run(
			func() string {
				return "IsValidPosition(" + strconv.Itoa(tt.x) + "," + strconv.Itoa(tt.y) + ")"
			}(),
			func(t *testing.T) {
				actual := IsValidPosition(tt.x, tt.y)
				if actual != tt.expected {
					t.Errorf("expected %v; got %v", tt.expected, actual)
				}
			},
		)
	}
}
