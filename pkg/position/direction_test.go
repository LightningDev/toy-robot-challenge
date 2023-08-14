package position

import "testing"

func TestDirectionString(t *testing.T) {
	tests := []struct {
		dir      Direction
		expected string
	}{
		{NORTH, "NORTH"},
		{SOUTH, "SOUTH"},
		{EAST, "EAST"},
		{WEST, "WEST"},
		{-5, ""},
	}

	for _, tt := range tests {
		got := tt.dir.String()
		if got != tt.expected {
			t.Errorf("Expected %s, got %s", tt.expected, got)
		}
	}
}

func TestDegree(t *testing.T) {
	tests := []struct {
		dir      Direction
		expected int
	}{
		{NORTH, 0},
		{SOUTH, 180},
		{EAST, 90},
		{WEST, 270},
	}

	for _, tt := range tests {
		got := tt.dir.Degree()
		if got != tt.expected {
			t.Errorf("Expected %d, got %d", tt.expected, got)
		}
	}
}

func TestStrToDirection(t *testing.T) {
	tests := []struct {
		input    string
		expected Direction
	}{
		{"NORTH", NORTH},
		{"North", NORTH},
		{"SOUTH", SOUTH},
		{"soutH", SOUTH},
		{"EAST", EAST},
		{"EAst", EAST},
		{"WEST", WEST},
		{"wEst", WEST},
		{"unknown", -1},
	}

	for _, tt := range tests {
		got := StrToDirection(tt.input)
		if got != tt.expected {
			t.Errorf("For input %s, expected %v, got %v", tt.input, tt.expected, got)
		}
	}
}

func TestDegreeToDirection(t *testing.T) {
	tests := []struct {
		input    int
		expected Direction
	}{
		{0, NORTH},
		{180, SOUTH},
		{90, EAST},
		{270, WEST},
		{45, -1},
	}

	for _, tt := range tests {
		got := DegreeToDirection(tt.input)
		if got != tt.expected {
			t.Errorf("For degree %d, expected %v, got %v", tt.input, tt.expected, got)
		}
	}
}
