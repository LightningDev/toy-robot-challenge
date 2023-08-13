package parser

import (
	"testing"

	"github.com/LightningDev/toy-robot-challenge/pkg/robot"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		input    string
		active   bool
		expected string
		hasError bool
	}{
		{"MOVE", false, "", true},                  // Should fail when robot isn't active
		{"PLACE 1,2,NORTH", false, "PLACE", false}, // Should work regardless of robot state
		{"MOVE", true, "MOVE", false},              // Should succeed when robot is active
		{"INVALIDCOMMAND", true, "", true},         // Invalid command
		{"", true, "", true},                       // Empty command
	}

	for _, tt := range tests {
		r := robot.Robot{Active: tt.active}
		cmd, err := ParseCommand(tt.input, r)

		if (err != nil) != tt.hasError {
			t.Errorf("Expected error: %v, but got: %v", tt.hasError, err)
			continue
		}

		if err == nil && cmd == nil {
			t.Errorf("Expected a command but got nil")
			continue
		}

		if cmd != nil && cmd.GetName() != tt.expected {
			t.Errorf("Expected command name to be %v, but got %v", tt.expected, cmd.GetName())
		}
	}
}
